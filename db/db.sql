CREATE SCHEMA `el-dorado`
USE `el-dorado`

CREATE TABLE `el-dorado`.`email` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`, `email`));

