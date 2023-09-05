package sqlstorage

const createTablesSQL = `
CREATE TABLE IF NOT EXISTS Users
(
	id      SERIAL PRIMARY KEY,
	login 	VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS TextData
(
    id SERIAL PRIMARY KEY,
    text_data TEXT,
    meta_info TEXT,
	user_id INTEGER REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS BankData
(
    id SERIAL PRIMARY KEY,
    card_number VARCHAR(16),
    card_expiry VARCHAR(5),
    CVV VARCHAR(3),
    meta_info TEXT,
	user_id INTEGER REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS BinaryData
(
    id SERIAL PRIMARY KEY,
    binary_data BYTEA,
    meta_info TEXT,
	user_id INTEGER REFERENCES Users(id)
);
`
