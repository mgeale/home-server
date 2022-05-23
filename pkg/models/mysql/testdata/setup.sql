CREATE TABLE balances (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    balance INTEGER NOT NULL,
    balanceAud INTEGER NOT NULL,
    pricebook INTEGER NOT NULL,
    product INTEGER NOT NULL,
    created DATETIME NOT NULL
);

CREATE INDEX idx_balances_created ON balances(created);

INSERT INTO balances (title, balance, balanceAud, pricebook, product, created) VALUES (
    'BAL-0022',
    100,
    1000,
    3333,
    2222,
    '2018-12-23 17:25:22'
);
