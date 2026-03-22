CREATE TABLE users (
                       id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       username VARCHAR(50) NOT NULL,
                       password CHAR(32) NOT NULL,
                       email VARCHAR(100),
                       created_at DATETIME,
                       updated_at DATETIME
);