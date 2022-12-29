CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR,
    password VARCHAR,
    created_at BIGINT,
    deleted_at BIGINT
);