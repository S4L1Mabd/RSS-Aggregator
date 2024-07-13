-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_time TIMESTAMP NOT NULL,
    update_time TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down

DROP TABLE users;

/*this file specialised for goose because role of goose if migrate a databse */
