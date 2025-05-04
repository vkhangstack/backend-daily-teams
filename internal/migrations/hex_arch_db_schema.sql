CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id         VARCHAR(18) PRIMARY KEY,
    uid        varchar(255)        not null,
    username   VARCHAR(50) UNIQUE  NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    first_name text,
    last_name  text,
    address    text,
    status     int8,
    avatar_url varchar(255),
    phone      varchar(14),
    provider_id varchar(14),
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by varchar(32),
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by varchar(32)
);

CREATE TABLE categories
(
    id      VARCHAR(18) PRIMARY KEY,
    user_id VARCHAR(18) REFERENCES users (id) ON DELETE CASCADE,
    name    VARCHAR(100)                                      NOT NULL,
    type    VARCHAR(10) CHECK (type IN ('income', 'expense')) NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by varchar(32),
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by varchar(32)
);

CREATE TABLE wallets
(
    id      VARCHAR(18) PRIMARY KEY,
    user_id VARCHAR(18) REFERENCES users (id) ON DELETE CASCADE,
    name    VARCHAR(100) NOT NULL,
    balance DECIMAL(15, 2) DEFAULT 0.00
);

CREATE TABLE transactions
(
    id               VARCHAR(18) PRIMARY KEY,
    user_id          VARCHAR(18) REFERENCES users (id) ON DELETE CASCADE,
    category_id      VARCHAR(18)                                                 DEFAULT NULL,
    amount           DECIMAL(15, 2)                                                NOT NULL,
    description      TEXT,
    transaction_type VARCHAR(10) CHECK (transaction_type IN ('income', 'expense')) NOT NULL,
    transaction_date TIMESTAMP DEFAULT NOW(),
    deleted_at       TIMESTAMP DEFAULT NULL,
    created_at       TIMESTAMP DEFAULT NOW(),
    created_by       varchar(32),
    updated_at       TIMESTAMP DEFAULT NOW(),
    updated_by       varchar(32)
);

CREATE TABLE budgets
(
    id          VARCHAR(18) PRIMARY KEY,
    user_id     VARCHAR(18) REFERENCES users (id) ON DELETE CASCADE,
    category_id VARCHAR(18) REFERENCES categories (id) ON DELETE CASCADE,
    amount      DECIMAL(15, 2) NOT NULL,
    start_date  DATE           NOT NULL,
    end_date    DATE           NOT NULL
);
