CREATE TYPE TRANSACTION_TYPE AS ENUM ('payment', 'refund');

CREATE TABLE IF NOT EXISTS users(
    ID UUID PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS merchants(
    ID UUID PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS cards(
    ID UUID PRIMARY KEY,
    owner UUID REFERENCES users(ID),
    account_balance DOUBLE PRECISION,
    blocked_amount DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS transactions(
    ID UUID PRIMARY KEY,
    sender UUID,
    receiver UUID,
    amount DOUBLE PRECISION,
    date DATE,
    type TRANSACTION_TYPE
);

CREATE TABLE IF NOT EXISTS authorizations(
    ID UUID PRIMARY KEY,
    merchant UUID REFERENCES merchants(ID),
    card UUID REFERENCES cards(ID),
    amount DOUBLE PRECISION,
    approved BOOLEAN,
    reversed DOUBLE PRECISION
);