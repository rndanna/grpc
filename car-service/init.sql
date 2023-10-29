CREATE SCHEMA cars;

CREATE TABLE cars
(
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL,
	price BIGINT,
    description TEXT,
	brand VARCHAR ( 50 ) NOT NULL,
	user_id INT NOT NULL,
	engine_id INT NOT NULL
);


INSERT INTO cars (name, price, description, brand, user_id, engine_id) VALUES ('car1', 123, 'description1', 'brand1', 1, 1);
INSERT INTO cars (name, price, description, brand, user_id, engine_id) VALUES ('car2', 1321, 'description2', 'brand1', 1, 2);
INSERT INTO cars (name, price, description, brand, user_id, engine_id) VALUES ('car3', 1321, 'description3', 'brand2', 2, 3);
