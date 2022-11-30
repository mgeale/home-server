CREATE TABLE IF NOT EXISTS balances (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    balance DECIMAL(16,9) NOT NULL,
    balanceaud DECIMAL(18,2),
    pricebookid INTEGER NOT NULL,
    productid INTEGER NOT NULL,
    created DATETIME NOT NULL
);

CREATE INDEX idx_balances_created ON balances(created);