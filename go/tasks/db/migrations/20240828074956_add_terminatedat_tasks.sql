-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';
alter table tasks add column terminatedAt timestamp default (datetime(current_timestamp, '+7 days'));
update tasks set terminatedAt = datetime(createdAt, '+7 days');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
alter table tasks drop column terminatedAt;
-- +goose StatementEnd
