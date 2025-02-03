-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
create table if not exists posts (
id bigserial primary key,
content text not null,
title text not null,
user_id bigint not null,
created_at timestamp(0) with time zone not null default now(),
updated_at timestamp(0) with time zone not null default now(),
tags varchar(100) []
-- constraint fk_user foreign key (user_id) references users(id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
drop table if exists posts;
-- +goose StatementEnd
