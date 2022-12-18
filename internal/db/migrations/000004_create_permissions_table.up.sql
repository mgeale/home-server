CREATE TABLE IF NOT EXISTS permissions (
	id   INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	code char(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
	user_id       INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
	permission_id INTEGER NOT NULL REFERENCES permissions ON DELETE CASCADE,
	PRIMARY KEY (user_id, permission_id)
);

INSERT INTO permissions (code) VALUES ('balances:read'), ('balances:write');