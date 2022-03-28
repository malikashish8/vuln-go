-- init db
CREATE DATABASE IF NOT EXISTS vulngo;
USE vulngo;
CREATE USER IF NOT EXISTS 'user'@'localhost';
GRANT ALL PRIVILEGES ON vulngo.* TO 'user'@'localhost';

-- create
CREATE TABLE ALBUM (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  title TEXT NOT NULL,
  artist TEXT NOT NULL,
  price FLOAT NOT NULL
) AUTO_INCREMENT=1;

-- insert
INSERT INTO ALBUM(title, artist, price) VALUES("Blue Train", "John Coltrane", 56.99);
INSERT INTO ALBUM(title, artist, price) VALUES("Jeru", "Gerry Mulligan", 17.99);
INSERT INTO ALBUM(title, artist, price) VALUES("Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99);
