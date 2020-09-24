# Таблица и данные

```sql
CREATE TABLE people (
	id SERIAL PRIMARY KEY,
	name varchar(255) NOT NULL,
	surname varchar(255),
	birth_date DATE,
	growth real,
	weight real,
	eyes varchar(255),
	hair varchar(255)
);

INSERT INTO people (name, surname, birth_date, growth, weight, eyes, hair)
VALUES ('ivan', 'ivanov', '11.03.1989', 180.3, 81.6, 'brown', 'brown'),
('ivan', 'petrov', '08.14.1991', 190.3, 81.6, 'blue', 'brown'),
('alexei', 'orlov', '03.26.1995', 187.3, 89.2, 'blue', 'blond'),
('maria', 'orlova', '08.28.1995', 164.3, 47.2, 'brown', 'blond'),
('alexandra', 'orlova', '11.11.1995', 167.3, 48.2, 'brown', 'blond');
```
