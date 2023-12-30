CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    username VARCHAR(255) UNIQUE,
    fullname VARCHAR(255),
    dob DATE,
    bio VARCHAR(255),
);

CREATE TABLE IF NOT EXISTS passwords (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    password VARCHAR(255),
    salt VARCHAR(255),
);