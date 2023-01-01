CREATE TABLE links (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    original_url VARCHAR,
    short_url VARCHAR,
    created_at BIGINT,
    deleted_at BIGINT
);