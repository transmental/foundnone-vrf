CREATE TABLE IF NOT EXISTS kms_eth_keys (
    id SERIAL PRIMARY KEY,
    address VARCHAR(42) NOT NULL UNIQUE,
    kms_key_id VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    ciphertext BYTEA NOT NULL,
    commitment TEXT NOT NULL,
    secret TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_kms_eth_keys_address ON kms_eth_keys(address);