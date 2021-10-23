BEGIN;

-- CREATE EXTENSION "uuid-ossp"

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  login varchar(100) NOT NULL UNIQUE,
  password varchar(5000) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE todo_lists (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title varchar(1000),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_by UUID NOT NULL REFERENCES users(id)
);

CREATE TABLE todos (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title varchar(1000),
  description varchar(5000),
  checked boolean NOT NULL DEFAULT false,
  todo_lists_id UUID NOT NULL REFERENCES todo_lists(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_by UUID NOT NULL REFERENCES users(id),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_by UUID NOT NULL REFERENCES users(id),
  deleted_at TIMESTAMPTZ
);

CREATE TABLE todo_changes (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title varchar(1000),
  description varchar(5000),
  checked boolean,
  todso_id UUID NOT NULL REFERENCES todos(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_by UUID NOT NULL REFERENCES users(id),
  deleted_at TIMESTAMPTZ
);

CREATE TABLE user_rights (
  users_id UUID REFERENCES users(id),
  todo_lists_id UUID REFERENCES todo_lists(id),
  rights varchar(10),
  PRIMARY KEY(users_id, todo_lists_id)
);

COMMIT;