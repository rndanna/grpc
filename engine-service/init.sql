CREATE SCHEMA engines;

CREATE TABLE engines
(
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL,
    description TEXT
);


INSERT INTO engines (name,  description) VALUES ('engine1', 'description1');
INSERT INTO engines (name,  description) VALUES ('engine2', 'description2');
INSERT INTO engines (name,  description) VALUES ('engine3', 'description3');

