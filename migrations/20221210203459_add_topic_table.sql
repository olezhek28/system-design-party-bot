-- +goose Up
create table topic
(
    id          bigint    not null,
    unit_id     bigint references unit (id),
    name        text      not null,
    description text      not null,
    link        text      not null,
    created_at  timestamp not null default now(),
    updated_at  timestamp
);

-- +goose Down
drop table topic;
