-- +migrate Up
CREATE TABLE accounts (
    id VARCHAR(32) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    user_name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE accounts;
