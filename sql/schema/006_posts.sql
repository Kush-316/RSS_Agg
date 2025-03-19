-- +goose Up

CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title text not null,
    description text,
    published_at TIMESTAMP not null,
    url text not null UNIQUE,
    feed_id UUID not null REFERENCES feeds(id) on delete CASCADE
);
-- +goose Down
DROP TABLE posts;