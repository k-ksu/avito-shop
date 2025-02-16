package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app" envPrefix:"APP_"`
		HTTP `yaml:"http"`
		PG   `yaml:"postgres" envPrefix:"PG_"`
	}

	// App -.
	App struct {
		Name       string        `yaml:"name"`
		Version    string        `yaml:"version"`
		TokenTTL   time.Duration `yaml:"token_ttl"`
		SigningKey string        `env:"SIGNING_KEY"`
	}

	// HTTP -.
	HTTP struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		SwaggerAddr string `yaml:"swagger_addr"`
	}

	// PG -.
	PG struct {
		PoolMax  int    `yaml:"pool_max"`
		URL      string `env:"URL"`
		URLLocal string `env:"URL_LOCAL"`
	}
)

// New returns app config.
func New() (*Config, error) {
	cfg := &Config{}

	pr, err := projectRootPath()
	if err != nil {
		return nil, fmt.Errorf("projectRootPath: %w", err)
	}

	if err = cfg.readConfigs(pr); err != nil {
		return nil, fmt.Errorf("cfg.readConfigs: %w", err)
	}

	if err = cfg.readSecrets(pr); err != nil {
		return nil, fmt.Errorf("cfg.readSecrets: %w", err)
	}

	return cfg, nil
}

func (cfg *Config) readConfigs(pr string) error {
	file, err := os.Open("/" + pr + "/config/config.yaml")
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("io.ReadAll: %w", err)
	}

	if err = yaml.Unmarshal(b, cfg); err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	return nil
}

func (cfg *Config) readSecrets(pr string) error {
	if err := godotenv.Load("/" + pr + "/.env"); err != nil {
		return fmt.Errorf("godotenv.Load: %w", err)
	}

	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("env.Parse: %w", err)
	}

	return nil
}

func projectRootPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("os.Getwd: %w", err)
	}

	parts := strings.Split(wd, "/")
	for i, part := range parts {
		if part == "avito-shop" {
			parts = parts[:i+1]
			break
		}
	}

	return path.Join(parts...), nil
}
