-- init db
CREATE DATABASE db;
USE db;
CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON db.* TO 'user'@'localhost';

-- create
CREATE TABLE ALBUM (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  title TEXT NOT NULL,
  artist TEXT NOT NULL,
  price FLOAT NOT NULL
);

-- insert
INSERT INTO ALBUM(title, artist, price) VALUES("Blue Train", "John Coltrane", 56.99);
INSERT INTO ALBUM(title, artist, price) VALUES("Jeru", "Gerry Mulligan", 17.99);
INSERT INTO ALBUM(title, artist, price) VALUES("Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99);
