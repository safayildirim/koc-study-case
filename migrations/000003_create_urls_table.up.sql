CREATE TABLE IF NOT EXISTS urls(
    id integer PRIMARY KEY,
    email VARCHAR (250) NOT NULL,
    original VARCHAR (250) NOT NULL,
    shortened_url VARCHAR (50) NOT NULL
);