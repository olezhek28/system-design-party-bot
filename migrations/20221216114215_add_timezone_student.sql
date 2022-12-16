-- +goose Up
alter table student add timezone integer;

-- +goose Down
alter table student drop timezone;
