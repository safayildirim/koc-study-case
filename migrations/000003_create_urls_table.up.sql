CREATE TABLE IF NOT EXISTS urls(
    id SERIAL PRIMARY KEY,
    email VARCHAR (250) NOT NULL,
    original VARCHAR (250) NOT NULL,
    shortened_id integer NOT NULL
);