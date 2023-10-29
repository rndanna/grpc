CREATE SCHEMA users;

CREATE TABLE users.users
(
	id SERIAL PRIMARY KEY,
	username VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL
);

INSERT INTO users.users (username, email) VALUES ('name1', 'email1@example.com');
INSERT INTO users.users (username, email) VALUES ('name2', 'email2@example.com');
