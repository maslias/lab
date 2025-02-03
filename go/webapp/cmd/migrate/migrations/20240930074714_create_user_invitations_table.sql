-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists user_invitations (
    token bytea primary key,
    user_id bigint not null,
    expire timestamp(0) with time zone not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
drop table if exists user_invitations;
-- +goose StatementEnd
