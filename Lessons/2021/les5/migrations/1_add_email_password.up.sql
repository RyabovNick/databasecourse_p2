BEGIN;

ALTER TABLE client
	ADD COLUMN email varchar(255);
	
ALTER TABLE client
	ADD COLUMN password varchar(255);

COMMIT;