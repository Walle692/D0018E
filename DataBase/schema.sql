CREATE SCHEMA IF NOT EXISTS myschema;

DROP TABLE IF EXISTS myschema.users;
DROP TABLE IF EXISTS myschema.token;

CREATE TABLE IF NOT EXISTS myschema.users (
    username VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
    password VARCHAR(50) NOT NULL,
    role VARCHAR(10) NOT NULL,
    token VARCHAR(256),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS myschema.token (
    username VARCHAR(50) REFERENCES myschema.users(username) PRIMARY KEY,
    token    VARCHAR(255) NOT NULL
);
