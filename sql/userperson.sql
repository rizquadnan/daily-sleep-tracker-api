CREATE TABLE userperson (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,                                                                                                                                           
  email VARCHAR(100) NOT NULL, 
  password_hash text NOT NULL
);

INSERT INTO userperson (name, email, password_hash)
VALUES ('test1', 'test1@gmail.com', 'somehashtest1'), ('test2', 'test2@gmail.com', 'somehashtest2'), ('test3', 'test3@gmail.com', 'somehashtest3');