INSERT INTO users(login, password) VALUES
  ('John', md5('12345')),
  ('Test', md5('54321'));

INSERT INTO user_tokens(user_id, token) VALUES
  (1, '6b9b9319'),
  (2, '9319aaf4');