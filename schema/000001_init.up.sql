CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE timestamp_template (
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

CREATE TABLE users (
  id serial primary key,
  email varchar(255) not null unique,
  password_hash varchar(255) not null
) INHERITS (timestamp_template);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();