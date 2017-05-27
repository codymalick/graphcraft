CREATE DATABASE graphcraft;

USE graphcraft;

CREATE TABLE item (
  id LONG PRIMARY KEY,
  description VARCHAR(50),
  name VARCHAR(50),
  icon VARCHAR(50),
  stackable INT
);

CREATE TABLE listing (
  id LONG PRIMARY KEY,
  item_id LONG,
  user_id LONG,
  bid LONG,
  buyout LONG,
  quantity INT,
  time LONG
)