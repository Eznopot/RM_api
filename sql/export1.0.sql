-- MariaDB dump 10.19  Distrib 10.9.2-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: db_RMS
-- ------------------------------------------------------
-- Server version	10.9.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `CV`
--

DROP TABLE IF EXISTS `CV`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CV` (
  `candidat_id` int(11) DEFAULT NULL,
  `competence` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `experience` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `formation` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `path` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  KEY `CV_FK` (`candidat_id`),
  KEY `CV_FK_1` (`user_id`),
  CONSTRAINT `CV_FK` FOREIGN KEY (`candidat_id`) REFERENCES `Candidat` (`id`),
  CONSTRAINT `CV_FK_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `CV`
--

LOCK TABLES `CV` WRITE;
/*!40000 ALTER TABLE `CV` DISABLE KEYS */;
INSERT INTO `CV` VALUES
(1,'[{insert:test plein de competence\n}]','[{insert: plein d\'experience\n}]','[{insert: plein de formation\n}]',NULL,NULL),
(2,'[{insert: autre test\n}]','[{insert: tets autre\n}]','[{insert: autres tets\n}]',NULL,NULL),
(3,'[{insert: Developper mobile\n}]','[{insert: Abeeway\n}]','[{insert: Epitech\n}]',NULL,NULL);
/*!40000 ALTER TABLE `CV` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Calendar`
--

DROP TABLE IF EXISTS `Calendar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Calendar` (
  `user_id` int(11) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `event_type` int(11) DEFAULT NULL,
  KEY `Calendar_FK` (`user_id`),
  CONSTRAINT `Calendar_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Calendar`
--

LOCK TABLES `Calendar` WRITE;
/*!40000 ALTER TABLE `Calendar` DISABLE KEYS */;
/*!40000 ALTER TABLE `Calendar` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Candidat`
--

DROP TABLE IF EXISTS `Candidat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Candidat` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `lastname` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Candidat`
--

LOCK TABLES `Candidat` WRITE;
/*!40000 ALTER TABLE `Candidat` DISABLE KEYS */;
INSERT INTO `Candidat` VALUES
(1,'Remy','test','remytest@gmail'),
(2,'dzqd','dzq','dzq'),
(3,'reel','reel','rell@reel');
/*!40000 ALTER TABLE `Candidat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `RDV`
--

DROP TABLE IF EXISTS `RDV`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `RDV` (
  `user_id` int(11) DEFAULT NULL,
  `candidat_id` int(11) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `appreciation` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  KEY `RDV_FK` (`user_id`),
  KEY `RDV_FK_1` (`candidat_id`),
  CONSTRAINT `RDV_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`),
  CONSTRAINT `RDV_FK_1` FOREIGN KEY (`candidat_id`) REFERENCES `Candidat` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RDV`
--

LOCK TABLES `RDV` WRITE;
/*!40000 ALTER TABLE `RDV` DISABLE KEYS */;
/*!40000 ALTER TABLE `RDV` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Token`
--

DROP TABLE IF EXISTS `Token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Token` (
  `user_id` int(11) DEFAULT NULL,
  `expiration` date DEFAULT NULL,
  `uuid` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  KEY `Token_FK` (`user_id`),
  CONSTRAINT `Token_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Token`
--

LOCK TABLES `Token` WRITE;
/*!40000 ALTER TABLE `Token` DISABLE KEYS */;
INSERT INTO `Token` VALUES
(7,NULL,'c978bff1-9c48-412d-a2ea-cfa714e662cc');
/*!40000 ALTER TABLE `Token` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `role` enum('user','owner','admin') COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES
(7,'Remy','098f6bcd4621d373cade4e832627b4f6','admin','remysalem@hotmail.fr');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Vacation`
--

DROP TABLE IF EXISTS `Vacation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Vacation` (
  `user_id` int(11) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  KEY `Vacation_FK` (`user_id`),
  CONSTRAINT `Vacation_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vacation`
--

LOCK TABLES `Vacation` WRITE;
/*!40000 ALTER TABLE `Vacation` DISABLE KEYS */;
/*!40000 ALTER TABLE `Vacation` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-26 12:03:22
