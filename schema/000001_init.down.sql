DROP TRIGGER IF EXISTS set_timestamp ON users;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS timestamp_template;

DROP FUNCTION IF EXISTS trigger_set_timestamp();