-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id serial primary key,
    name text,
    obfuscated_password text,
    coins bigint default 1000
);

create unique index user_name_idx on users(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
