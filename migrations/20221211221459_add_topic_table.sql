-- +goose Up
create table topic
(
    id         bigserial primary key,
    name       text      not null,
    link       text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table topic;
