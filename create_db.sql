-- create_db.sql
drop database library;
-- Création de la base de données "library"
CREATE DATABASE IF NOT EXISTS library;
USE library;

-- Table pour les livres
CREATE TABLE IF NOT EXISTS books (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  item_type VARCHAR(255),
  status VARCHAR(255),
  created_at DATETIME,
  updated_at DATETIME
);

-- Table pour les emprunts
CREATE TABLE IF NOT EXISTS loans (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  resource_id INT NOT NULL,
  loan_date DATETIME,
  return_date DATETIME,
  status VARCHAR(255)
);

-- Table pour les utilisateurs
CREATE TABLE IF NOT EXISTS users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);
