create table note
(
    id         bigserial primary key,
    name       varchar not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);