CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR (250) UNIQUE NOT NULL,
    password VARCHAR (250) NOT NULL,
    subscription_type integer NOT NULL
);