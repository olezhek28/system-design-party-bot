-- +goose Up
create table student
(
    id          bigserial primary key,
    name        text      not null,
    telegram_id integer   not null,
    created_at  timestamp not null default now()
);

-- +goose Down
drop table student;
