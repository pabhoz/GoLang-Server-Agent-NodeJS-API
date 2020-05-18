-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema servers
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `servers` ;

-- -----------------------------------------------------
-- Schema servers
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `servers` DEFAULT CHARACTER SET utf8 ;
USE `servers` ;

-- -----------------------------------------------------
-- Table `servers`.`Agents`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `servers`.`Agents` (
  `uid` VARCHAR(60) NOT NULL,
  `createdAt` TIMESTAMP NOT NULL,
  PRIMARY KEY (`uid`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `servers`.`ProcessorLogs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `servers`.`ProcessorLogs` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `agentId` VARCHAR(60) NOT NULL,
  `cpuIndex` INT NOT NULL,
  `vendorId` VARCHAR(45) NOT NULL,
  `family` VARCHAR(45) NOT NULL,
  `numberOfCores` INT NOT NULL,
  `modelName` VARCHAR(45) NOT NULL,
  `speed` VARCHAR(45) NOT NULL,
  `currentCPUUtilization` TEXT NOT NULL,
  `createdAt` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_ProcessorLogs_Agents1`
    FOREIGN KEY (`agentId`)
    REFERENCES `servers`.`Agents` (`uid`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_ProcessorLogs_Agents1_idx` ON `servers`.`ProcessorLogs` (`agentId` ASC);


-- -----------------------------------------------------
-- Table `servers`.`UsersLogs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `servers`.`UsersLogs` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `agentId` VARCHAR(60) NOT NULL,
  `activeUsers` TEXT NOT NULL,
  `createdAt` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_UsersLogs_Agents1`
    FOREIGN KEY (`agentId`)
    REFERENCES `servers`.`Agents` (`uid`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_UsersLogs_Agents1_idx` ON `servers`.`UsersLogs` (`agentId` ASC);


-- -----------------------------------------------------
-- Table `servers`.`RunningProcessesLog`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `servers`.`RunningProcessesLog` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `agentId` VARCHAR(60) NOT NULL,
  `total` VARCHAR(45) NOT NULL,
  `running` VARCHAR(45) NOT NULL,
  `processesList` TEXT NOT NULL,
  `createdAt` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_RunningProcessesLog_Agents1`
    FOREIGN KEY (`agentId`)
    REFERENCES `servers`.`Agents` (`uid`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_RunningProcessesLog_Agents1_idx` ON `servers`.`RunningProcessesLog` (`agentId` ASC);


-- -----------------------------------------------------
-- Table `servers`.`SOLogs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `servers`.`SOLogs` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `agentId` VARCHAR(60) NOT NULL,
  `runtime` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `platform` VARCHAR(45) NOT NULL,
  `createdAt` TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_SOLogs_Agents1`
    FOREIGN KEY (`agentId`)
    REFERENCES `servers`.`Agents` (`uid`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_SOLogs_Agents1_idx` ON `servers`.`SOLogs` (`agentId` ASC);


-- -----------------------------------------------------
-- Table `servers`.`ActiveUsers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `servers`.`ActiveUsers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `ussersLogId` INT NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `application` VARCHAR(45) NOT NULL,
  `date` VARCHAR(45) NOT NULL,
  `time` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_ActiveUsers_UsersLogs1`
    FOREIGN KEY (`ussersLogId`)
    REFERENCES `servers`.`UsersLogs` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE INDEX `fk_ActiveUsers_UsersLogs1_idx` ON `servers`.`ActiveUsers` (`ussersLogId` ASC);


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
