-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
create table tasks (
    id integer primary key autoincrement,
    title varchar(255) not null,
    details varchar(255),
    createdAt datetime default current_timestamp,
    doneAt datetime
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tasks;
-- SELECT 'down SQL query';
-- +goose StatementEnd
