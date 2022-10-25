CREATE TABLE balances (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    balance DECIMAL(16,9) NOT NULL,
    balanceaud DECIMAL(18,2),
    pricebookid INTEGER NOT NULL,
    productid INTEGER NOT NULL,
    created DATETIME NOT NULL
);

CREATE INDEX idx_balances_created ON balances(created);

INSERT INTO balances (name, balance, balanceaud, pricebookid, productid, created) VALUES (
    'BAL-0022',
    100.89,
    1000.01,
    3333,
    2222,
    '2018-12-23 17:25:22'
);

CREATE TABLE transactions (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    amount DECIMAL(9,2) NOT NULL,
    date VARCHAR(100) NOT NULL,
    type VARCHAR(100) NOT NULL,
    created DATETIME NOT NULL
);

INSERT INTO transactions (name, amount, date, type, created) VALUES (
    'name',
    100,
    '2018-12-23 17:25:22',
    'Repayment',
    '2018-12-23 17:25:22'
);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (name, email, hashed_password, created) VALUES (
    'Alice Jones',
    'alice@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2018-12-23 17:25:22'
);