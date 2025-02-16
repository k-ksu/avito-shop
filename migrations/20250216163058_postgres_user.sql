-- +goose Up
-- +goose StatementBegin
CREATE USER postgres WITH SUPERUSER PASSWORD 'password';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
