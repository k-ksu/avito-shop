-- +goose Up
-- +goose StatementBegin
create table if not exists transaction_history (
     from_user int references users(id),
     to_user   int references users(id),
     amount   int,
     created_at timestamp default now()
);

create index transaction_history_from_user_idx on transaction_history(from_user);
create index transaction_history_to_user_idx on transaction_history(to_user);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists transaction_history
-- +goose StatementEnd
