-- +migrate Up
CREATE TABLE accounts (
    id VARCHAR(32) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- +migrate Down
DROP TABLE accounts;
