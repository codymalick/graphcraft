CREATE DATABASE graphcraft;

USE graphcraft;

CREATE TABLE items (
  id INT NOT NULL,
  description VARCHAR(250),
  name VARCHAR(50),
  icon VARCHAR(50),
  stackable INT,
  PRIMARY KEY (id)
);

CREATE TABLE users (
  id INT NOT NULL PRIMARY KEY,
  username VARCHAR (50) NOT NULL,
  server VARCHAR (50) NOT NULL
);

CREATE TABLE listings (
  id INT NOT NULL,
  item_id INT,
  user_id INT,
  bid LONG,
  buyout LONG,
  quantity INT,
  time LONG,
  PRIMARY KEY (id),
  FOREIGN KEY (item_id) REFERENCES item(id),
  FOREIGN KEY (user_id) REFERENCES user(id)
);