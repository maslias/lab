-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
create table if not exists comments (
id bigserial primary key,
post_id bigint not null,
user_id bigint not null,
content text not null,
created_at timestamp(0) with time zone not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
drop table if exists comments;
-- +goose StatementEnd
