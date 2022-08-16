CREATE TABLE customers
(
    id            SERIAL PRIMARY KEY,
    uid           UUID UNIQUE              NOT NULL DEFAULT gen_random_uuid(),
    name          VARCHAR(100)             NOT NULL,
    date_of_birth DATE                     NOT NULL,
    city          VARCHAR(100)             NOT NULL,
    zipcode       VARCHAR(10)              NOT NULL,
    status       SMALLINT                 NOT NULL DEFAULT 1,
    updated_by    VARCHAR(100)             NOT NULL,
    updated_date  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts
(
    id           SERIAL PRIMARY KEY,
    uid          UUID UNIQUE              NOT NULL DEFAULT gen_random_uuid(),
    customer_uid UUID                     NOT NULL REFERENCES customers (uid) ON DELETE RESTRICT,
    opening_date DATE                     NOT NULL,
    account_type VARCHAR(10)              NOT NULL,
    amount       NUMERIC(18, 2) CHECK (amount >= 0),
    status       SMALLINT                 NOT NULL DEFAULT 1,
    updated_by   VARCHAR(100)             NOT NULL,
    updated_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions
(
    id               SERIAL PRIMARY KEY,
    uid              UUID UNIQUE              NOT NULL DEFAULT gen_random_uuid(),
    account_uid      UUID                     NOT NULL REFERENCES accounts (uid) ON DELETE RESTRICT,
    amount           NUMERIC(18, 2) CHECK (amount >= 0),
    transaction_type VARCHAR(10)              NOT NULL,
    created_by       VARCHAR(100)             NOT NULL,
    transaction_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);