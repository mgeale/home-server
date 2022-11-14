CREATE TABLE IF NOT EXISTS transactions (
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