CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    role varchar(255) NOT NULL,
    avatar varchar(255) NULL
);