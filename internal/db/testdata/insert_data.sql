INSERT INTO users (id, name, email, hashed_password, created) VALUES (
    UUID(),
    'Alice Jones',
    'alice@example.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2018-12-23 17:25:22'
);

INSERT INTO balances (id, name, balance, balanceaud, pricebookid, productid, created) VALUES (
    "7a59f3e8-b0b9-11ed-a356-0242ac110002",
    'BAL-0022',
    100.89,
    1000.01,
    "01s9D000001lX8rQAE",
    "01t9D000003rsQoQAI",
    '2018-12-23 17:25:22'
), (
    "7a59f5c1-b0b9-11ed-a356-0242ac110002",
    'BAL-0033',
    85.12,
    52.78,
    "01s9D000001lX8rQAE",
    "01t9D000003rsQoQAI",
    '2018-12-23 17:25:22'
), (
    "7a59f61a-b0b9-11ed-a356-0242ac110002",
    'BAL-0033',
    45.12,
    47.73,
    "01s9D000001lX8rQAE",
    "01t9D000003rsQoQAI",
    '2018-12-23 17:25:22'
);

INSERT INTO transactions (id, name, amount, date, type, created) VALUES (
    "1c0d2b44-b0ce-11ed-b95f-dca632bb7cae",
    'name',
    100,
    '2018-12-23 17:25:22',
    'Repayment',
    '2018-12-23 17:25:22'
);
