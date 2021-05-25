-- MySQL dump 10.13  Distrib 8.0.16, for macos10.14 (x86_64)
--
-- Host: localhost    Database: cashlez
-- ------------------------------------------------------
-- Server version	5.7.26-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `reconcile_raw_details_at_mandiri`
--

DROP TABLE IF EXISTS `reconcile_raw_details_at_mandiri`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `reconcile_raw_details_at_mandiri` (
  `mid` varchar(50) NOT NULL,
  `merchant_official` varchar(100) NOT NULL,
  `trading_name` varchar(255) NOT NULL,
  `bank_mandiri_acc` varchar(50) NOT NULL,
  `other_bank_acc` varchar(50) DEFAULT NULL,
  `merchacc` int(11) DEFAULT NULL,
  `trxdate` date NOT NULL,
  `settledate` date NOT NULL,
  `trxcode` varchar(10) DEFAULT NULL,
  `description` varchar(100) DEFAULT NULL,
  `card` varchar(50) NOT NULL,
  `crdtype` varchar(10) NOT NULL,
  `tid` varchar(50) NOT NULL,
  `authcode` varchar(10) NOT NULL,
  `paymentbatch` varchar(50) DEFAULT NULL,
  `tidbatch` varchar(10) DEFAULT NULL,
  `batchseq` varchar(10) DEFAULT NULL,
  `amount` decimal(20,0) NOT NULL,
  `nonmdramount` decimal(20,0) DEFAULT NULL,
  PRIMARY KEY (`mid`,`trading_name`,`trxdate`,`tid`,`authcode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-24  1:43:25
