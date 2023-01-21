-- +migrate Up
CREATE TABLE webauthn_registrations (
    id VARCHAR(32) PRIMARY KEY,
    account_id VARCHAR(32) NOT NULL REFERENCES accounts ON UPDATE NO ACTION ON DELETE CASCADE,
    challenge VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX ON webauthn_registrations(account_id);

-- +migrate Down
DROP TABLE webauthn_credentials;
