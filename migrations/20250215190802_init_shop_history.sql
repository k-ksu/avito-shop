-- +goose Up
-- +goose StatementBegin
create table if not exists shop_history (
    id serial primary key,
    user_id int references users(id),
    item_id int references merch_items(id),
    created_at timestamp default now()
);

create index shop_history_user_id_idx on shop_history(user_id);
create index shop_history_item_id_idx on shop_history(item_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists shop_history;
-- +goose StatementEnd
