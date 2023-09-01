create table "user"
(
    id   bigserial primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);