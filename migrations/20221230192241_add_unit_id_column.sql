-- +goose Up
alter table meeting add unit_id bigint references unit (id);
alter table topic add unit_id bigint references unit (id);

-- +goose Down
alter table meeting drop unit_id;
alter table topic drop unit_id;
