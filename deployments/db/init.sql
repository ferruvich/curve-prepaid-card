CREATE TYPE TRANSACTION_TYPE AS ENUM ('payment', 'refund');
CREATE TYPE TRANSACTION_STATUS AS ENUM ('not_catched', 'partially_catched', 'completed');

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
    date TIMESTAMP,
    type TRANSACTION_TYPE,
    status TRANSACTION_STATUS
);

CREATE TABLE IF NOT EXISTS authorizations(
    ID UUID PRIMARY KEY,
    merchant UUID REFERENCES merchants(ID),
    card UUID REFERENCES cards(ID),
    amount DOUBLE PRECISION,
    reversed DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS captures(
    ID UUID PRIMARY KEY,
    transaction UUID REFERENCES transactions(ID),
    auth UUID REFERENCES authorizations(ID),
    amount DOUBLE PRECISION
);