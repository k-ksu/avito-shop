package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/k-ksu/avito-shop/pkg/postgres"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/pressly/goose/v3"
)

var (
	Host     = "localhost"
	Name     = "test"
	Password = "test"
	Port     = "5432"
	User     = "test"

	migrationDir = "../../migrations"

	testRepo *postgres.Client
)

// nolint:perfsprint
func NewTestDocker() (*postgres.Client, func()) {
	var (
		dsn      string
		resource *dockertest.Resource
		db       *pgxpool.Pool
		ctx      = context.Background()
	)

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal("cannot connect to test docker ", err)
	}

	resource, err = pool.Run("postgres", "13", []string{
		fmt.Sprintf("POSTGRES_USER=%s", User),
		fmt.Sprintf("POSTGRES_PASSWORD=%s", Password),
		fmt.Sprintf("POSTGRES_DB=%s", Name),
	})
	if err != nil {
		log.Fatal("pool run failed", err)
	}

	closeFunc := func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatal("cannot close resource", err)
		}
	}

	PortDB, _ := strconv.Atoi(resource.GetPort("5432/tcp"))
	err = pool.Retry(func() error {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, PortDB, User, Password, Name)
		db, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return errors.New("testDB connection failed: " + err.Error())
		}

		err = db.Ping(ctx)
		if err != nil {
			return errors.New("ping failed: " + err.Error())
		}

		dbSQL, oErr := sql.Open("postgres", dsn)
		if oErr != nil {
			return errors.New("dbSQL failed: " + err.Error())
		}

		defer dbSQL.Close()
		goose.SetTableName("goose_db_version")
		err = goose.SetDialect("postgres")
		if err != nil {
			return errors.New("set dialect failed: " + err.Error())
		}
		err = goose.Up(dbSQL, migrationDir)
		if err != nil {
			return errors.New("migration up failed: " + err.Error())
		}

		return nil
	})
	if err != nil {
		closeFunc()
		log.Fatal(err)
	}

	cl := &postgres.Client{
		Pool: db,
	}

	return cl, closeFunc
}

func TestMain(m *testing.M) {
	var exitCode int
	func() {
		db, closeFunc := NewTestDocker()
		defer closeFunc()

		testRepo = db
		exitCode = m.Run()
	}()

	os.Exit(exitCode)
}
