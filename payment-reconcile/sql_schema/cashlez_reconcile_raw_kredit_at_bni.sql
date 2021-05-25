-- MySQL dump 10.13  Distrib 8.0.16, for macos10.14 (x86_64)
--
-- Host: localhost    Database: cashlez
-- ------------------------------------------------------
-- Server version	5.7.26-log
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;
SET NAMES utf8;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */
;
/*!40103 SET TIME_ZONE='+00:00' */
;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */
;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */
;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */
;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */
;
--
-- Table structure for table `reconcile_raw_details_at_mandiri`
--
DROP TABLE IF EXISTS `cashlez_reconcile_raw_kredit_at_bni`;
/*!40101 SET @saved_cs_client     = @@character_set_client */
;
SET character_set_client = utf8mb4;
CREATE TABLE `cashlez_reconcile_raw_kredit_at_bni` (
    `proc_date` date NOT NULL,
    `mid` varchar(50) NOT NULL,
    `ob` varchar(50),
    `gb` varchar(50),
    `seq` VARCHAR(100),
    `type` varchar(10),
    `trxdate` date NOT NULL,
    `authcode` varchar(10) NOT NULL,
    `card_no` varchar(50) NOT NULL,
    `amount` decimal(20, 0) NOT NULL,
    `tid` VARCHAR(50),
    `trxtype` VARCHAR(10),
    `ptr` varchar(50) NOT NULL,
    `rate` DECIMAL(20, 0),
    `disc_amount` DECIMAL(20, 0),
    `air_fare` VARCHAR(50) NOT NULL,
    `plan` VARCHAR(50),
    `ss_amount` DECIMAL(20, 0) NOT NULL,
    `ss_fee` DECIMAL(20, 0) NOT NULL,
    `flag` VARCHAR(50),
    `nett_amount` DECIMAL(20, 0) NOT NULL,
    `merchant_account` VARCHAR(100) NOT NULL,
    `merchant_name` VARCHAR(100) NOT NULL,
    PRIMARY KEY (
        `mid`,
        `trxtype`,
        `trxdate`,
        `tid`,
        `authcode`
    )
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
/*!40101 SET character_set_client = @saved_cs_client */
;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */
;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */
;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */
;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */
;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */
;
-- Dump completed on 2020-06-24  1:43:25