CREATE TABLE IF NOT EXISTS users
(
    uuid       varchar(50) UNIQUE,
    username text,
    email    text UNIQUE,
    password text,
    isverified boolean
    );

CREATE TABLE IF NOT EXISTS verif
(
    email text UNIQUE,
    code text,
    isverified bool
    );