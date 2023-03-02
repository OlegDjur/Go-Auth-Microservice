CREATE TABLE users (
    id bigserial PRIMARY KEY,
    usernmae VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL
    password VARCHAR(250) NOT NULL 
);

