BEGIN

CREATE UNIQUE INDEX email_unique_idx ON client (email);

ALTER TABLE client
	ALTER COLUMN email SET NOT NULL;

ALTER TABLE client
	ALTER COLUMN password SET NOT NULL;

COMMIT;