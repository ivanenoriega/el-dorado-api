DROP SCHEMA `el-dorado`;
CREATE SCHEMA `el-dorado`;
USE `el-dorado`;

CREATE TABLE email (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`, `email`));

CREATE TABLE contact (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `phone` VARCHAR(45) NOT NULL,
  `subject` VARCHAR(45) NOT NULL,
  `message` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`));

