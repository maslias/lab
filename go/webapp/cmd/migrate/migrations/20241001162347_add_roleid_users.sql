-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
alter table if exists users
add column role_id int references roles(id) default 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists users
drop column role_id;
-- +goose StatementEnd
