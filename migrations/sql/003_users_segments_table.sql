CREATE TABLE users_segments (
    user_id INT REFERENCES users(id),
    slug VARCHAR(255) REFERENCES segments(slug),
    CONSTRAINT unique_user_segment UNIQUE (user_id, slug)
);

---- create above / drop below ----

DROP TABLE users_segments;