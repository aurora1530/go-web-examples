DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  password_hash bytea NOT NULL,
  username VARCHAR(255) NOT NULL UNIQUE
);
