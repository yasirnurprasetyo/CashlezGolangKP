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

DROP TABLE IF EXISTS `20200522_ecom_cashlez_bri`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `20200522_ecom_cashlez_bri` (
  `no` int(11) NOT NULL,
  `nama_merchant` varchar(100) NOT NULL,
  `rk_date` date NOT NULL,
  `proc_date` date NOT NULL,
  `mid` varchar(50) NOT NULL,
  `cardtype` date NOT NULL,
  `trx_date` date NOT NULL,
  `auth` varchar(10) NOT NULL,
  `cardno` varchar(20) NOT NULL,
  `jenis_trx` varchar(20) NOT NULL,
  `amount` int(11) NOT NULL,
  `nonfare` int(11) NOT NULL,
  `rate` decimal(20,0) NOT NULL,
  `disc_amt` decimal(20,0) DEFAULT NULL,
  `airfare` decimal(20,0) DEFAULT NULL,
  `flag` int(11) DEFAULT NULL,
  `net_amt` decimal(20,0) NOT NULL,
  `merchant_ref_number` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`auth`,`cardno`)
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
