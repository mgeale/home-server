CREATE TABLE balances (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    balance INTEGER NOT NULL,
    balanceaud INTEGER NOT NULL,
    pricebookid INTEGER NOT NULL,
    productid INTEGER NOT NULL,
    created DATETIME NOT NULL
);

CREATE INDEX idx_balances_created ON balances(created);

INSERT INTO balances (name, balance, balanceaud, pricebookid, productid, created) VALUES (
    'BAL-0022',
    100,
    1000,
    3333,
    2222,
    '2018-12-23 17:25:22'
);
