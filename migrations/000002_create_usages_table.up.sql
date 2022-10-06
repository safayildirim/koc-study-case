CREATE TABLE IF NOT EXISTS usages(
    email VARCHAR (250) UNIQUE NOT NULL,
    remaining integer NOT NULL
);