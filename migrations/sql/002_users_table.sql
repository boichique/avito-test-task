CREATE TABLE users (
    id INT PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

---- create above / drop below ----

DROP TABLE users;