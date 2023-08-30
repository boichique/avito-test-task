CREATE TABLE segments (
    slug VARCHAR(255) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

---- create above / drop below ----

DROP TABLE segments;
