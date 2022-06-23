CREATE TABLE menu (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	price money NOT NULL,
	description varchar(3000),
	weight INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO menu (name, price, weight) VALUES
('Пицца пепперони', 15, 700),
('Салат', 10, 3000),
('Шаурма', 2, 200);

CREATE TABLE client (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	address varchar(1000) NOT NULL,
	phone varchar(11) NOT NULL
);

INSERT INTO client (name, address, phone) VALUES
('Иван', 'Боголюбова, 21, кв 14', '88005553535'),
('Николай', 'Университетская, 2, кв 11', '83004441122'),
('Вера', 'Пушкина, 4, кв 27', '89153002222');

CREATE TABLE order_ (
	id SERIAL PRIMARY KEY,
	client_id INTEGER NOT NULL REFERENCES client(id),
	created_at TIMESTAMP NOT NULL DEFAULT now()
);

INSERT INTO order_ (client_id) VALUES 
(1),(1),(1),(2);

CREATE TABLE order_menu (
	order_id INTEGER NOT NULL REFERENCES order_(id),
	menu_id INTEGER NOT NULL REFERENCES menu(id),
	count INTEGER NOT NULL,
	price money NOT NULL,
	PRIMARY KEY(order_id, menu_id)
);

INSERT INTO order_menu (order_id, menu_id, count, price) VALUES
(1, 1, 1, 15),
(1, 2, 1, 10),
(2, 1, 2, 30),
(3, 3, 6, 12),
(4, 1, 2, 30);