-- +goose Up
create table meeting
(
    id          bigserial primary key,
    unit_id     bigint,
    topic_id    bigint,
    status      text      not null,
    start_date  timestamp not null,
    speaker_id  bigint references student (id),
    listener_id bigint references student (id),
    created_at  timestamp not null default now()
);

-- +goose Down
drop table meeting;
