-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
create extension if not exists citext;
create table if not exists users (
id bigserial primary key,
email citext unique not null,
username varchar(255) not null,
password bytea not null,
created_at timestamp(0) with time zone not null default now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
drop table if exists users;
-- +goose StatementEnd
