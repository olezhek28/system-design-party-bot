-- +goose Up
create table unit
(
    id          bigserial primary key,
    name        text      not null,
    description text      not null,
    link        text      not null,
    created_at  timestamp not null default now(),
    updated_at  timestamp
);

-- +goose Down
drop table unit;