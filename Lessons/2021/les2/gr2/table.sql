CREATE TABLE brand (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	country varchar(255) NOT NULL
);

CREATE TABLE car (
	id SERIAL PRIMARY KEY,
	brand_id INTEGER NOT NULL REFERENCES brand(id),
	model varchar(255) NOT NULL,
	cost money,
	year_of_creation NUMERIC(4,0) NOT NULL,
	is_available boolean NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NOT NULL DEFAULT now(),
	deleted_at TIMESTAMP
);

INSERT INTO brand (name, country) VALUES 
('BMW', 'Germany'),
('Lada', 'Russia'),
('Honda', 'Japan'),
('Tesla', 'USA');

INSERT INTO car (brand_id, model, cost, year_of_creation, is_available) VALUES
(1, '520D', 2000000, 2017, true),
(2, 'XRay', 800000, 2018, true),
(4, 'Model S', 12500000, 2020, false),
(1, 'M5', 7000000, 2018, true);

CREATE TABLE manager (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	surname varchar(255) NOT NULL,
	car_id INTEGER REFERENCES car(id)
);

INSERT INTO manager (name, surname, car_id) VALUES 
('Ivan', 'Ivanov', 1),
('Petr', 'Petrov', 2),
('Alex', 'Alexandrov', null);