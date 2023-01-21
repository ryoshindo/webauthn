-- +migrate Up
CREATE TABLE webauthn_credentials (
    id VARCHAR(32) PRIMARY KEY,
    account_id VARCHAR(32) NOT NULL REFERENCES accounts ON UPDATE NO ACTION ON DELETE CASCADE,
    public_key VARCHAR(255),
    external_id VARCHAR(255),
    sign_count BIGINT,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX ON webauthn_credentials(account_id);

-- +migrate Down
DROP TABLE webauthn_credentials;
