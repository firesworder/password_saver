package sqlstorage

const createTablesSQL = `
CREATE TABLE IF NOT EXISTS Users
(
	id      SERIAL PRIMARY KEY,
	login 	VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Records
(
    id SERIAL PRIMARY KEY,
    record_type VARCHAR(30),
    content BYTEA,
    meta_info TEXT,
    user_id INTEGER REFERENCES Users(id)
);
`
