-- +goose Up
-- +goose StatementBegin
create table if not exists merch_items (
    id     serial primary key,
    name   text,
    price  int
);

insert into merch_items (name, price) values
 ('t-shirt', 80),
 ('cup', 20),
 ('book', 50),
 ('pen', 10),
 ('powerbank', 230),
 ('hoody', 300),
 ('umbrella', 200),
 ('socks', 10),
 ('wallet', 50),
 ('pink-hoody', 500);


create unique index merch_items_name_idx on merch_items(name);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists merch_items
-- +goose StatementEnd
