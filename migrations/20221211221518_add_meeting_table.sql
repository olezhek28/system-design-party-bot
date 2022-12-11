-- +goose Up
create table meeting
(
    id          bigserial primary key,
    topic_id    bigint references topic (id),
    status text not null,
    start_date  timestamp not null,
    speaker_id  bigint references student (id),
    listener_id bigint references student (id),
    created_at  timestamp not null default now()
);

-- +goose Down
drop table meeting;
