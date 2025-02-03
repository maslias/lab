-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
alter table if exists users
add column is_active boolean not null default false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists users
drop column is_active;
-- +goose StatementEnd
