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
    content BYTEA,
    met_info TEXT,
    type VARCHAR(30),
    user_id INTEGER REFERENCES Users(id)
);
`
