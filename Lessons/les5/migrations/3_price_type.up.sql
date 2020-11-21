BEGIN;

ALTER TABLE menu
	ALTER COLUMN price TYPE numeric(8,2);

ALTER TABLE order_menu
	ALTER COLUMN price TYPE numeric(8,2);

COMMIT;