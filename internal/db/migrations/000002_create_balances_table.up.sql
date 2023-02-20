CREATE TABLE IF NOT EXISTS balances (
    id UUID NOT NULL,
    PRIMARY KEY(id),
    name VARCHAR(100) NOT NULL,
    balance DECIMAL(16,9) NOT NULL,
    balanceaud DECIMAL(18,2),
    pricebookid VARCHAR(18) NOT NULL,
    productid VARCHAR(18) NOT NULL,
    created DATETIME NOT NULL
);

CREATE INDEX idx_balances_created ON balances(created);