-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
create table if not exists roles (
id bigserial primary key,
name varchar(255) not null unique,
level int not null default 0,
description text
);

insert into roles (name, description, level)
values (
'user',
'user can do user stuff',
1
);

insert into roles (name, description, level)
values (
'moderator',
'moderator stuf',
2
);

insert into roles (name, description, level)
values (
'admin',
'playing god',
3
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
drop table if exists roles;
-- +goose StatementEnd
