CREATE DATABASE graphcraft;

USE graphcraft;

CREATE TABLE items (
  id INT NOT NULL,
  description VARCHAR(500),
  name VARCHAR(250),
  icon VARCHAR(250),
  stackable INT,
  PRIMARY KEY (id)
);

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  username VARCHAR (50) NOT NULL,
  server VARCHAR (50) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE timestamps (
  id INT NOT NULL AUTO_INCREMENT,
  timestamp long,
  PRIMARY KEY (id)
);

CREATE TABLE listings (
  id INT NOT NULL AUTO_INCREMENT,
  listing_id INT NOT NULL,
  item_id INT,
  user_id INT,
  b_id LONG,
  buyout LONG,
  quantity INT,
  timestamp_id INT,
  PRIMARY KEY (id),
  FOREIGN KEY (item_id) REFERENCES items(id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (timestamp_id) REFERENCES  timestamps(id)
);
