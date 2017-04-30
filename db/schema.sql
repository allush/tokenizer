CREATE TABLE users
(
  id       SERIAL PRIMARY KEY NOT NULL,
  login    VARCHAR,
  password VARCHAR,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_tokens
(
  id      SERIAL PRIMARY KEY NOT NULL,
  user_id INTEGER            NOT NULL,
  token   VARCHAR,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX index_user_tokens_on_user_id
  ON user_tokens (user_id);
CREATE UNIQUE INDEX index_user_tokens_on_token
  ON user_tokens (token);