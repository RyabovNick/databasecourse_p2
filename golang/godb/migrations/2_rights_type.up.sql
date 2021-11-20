BEGIN;

UPDATE user_rights
SET rights = '5'
WHERE rights = 'o';

ALTER TABLE user_rights
  ALTER rights TYPE INT USING rights::integer;

ALTER TABLE todo_changes
  RENAME COLUMN todso_id TO todos_id;

COMMIT;