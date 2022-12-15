-- +goose Up
create table student
(
    id                bigserial primary key,
    first_name        text      not null,
    last_name         text      not null,
    telegram_id       integer   not null,
    telegram_username text      not null,
    created_at        timestamp not null default now()
);

-- +goose Down
drop table student;
