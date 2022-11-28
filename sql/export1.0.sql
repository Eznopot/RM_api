-- MariaDB dump 10.19  Distrib 10.9.4-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: db_RMS
-- ------------------------------------------------------
-- Server version	10.9.4-MariaDB

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
-- Table structure for table `AdminInfo`
--

DROP TABLE IF EXISTS `AdminInfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `AdminInfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `value` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `AdminInfo`
--

LOCK TABLES `AdminInfo` WRITE;
/*!40000 ALTER TABLE `AdminInfo` DISABLE KEYS */;
INSERT INTO `AdminInfo` VALUES
(46,'Bonjour,\nLe site est en developpement, si vous rencontrer un bug ou un probleme merci de contacter un administrateur.\nN\'hesiter pas a proposer de changement ou de nouvelle fonctionalité.\nBonne journée !');
/*!40000 ALTER TABLE `AdminInfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `CV`
--

DROP TABLE IF EXISTS `CV`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `CV` (
  `candidat_id` int(11) DEFAULT NULL,
  `competence` text DEFAULT NULL,
  `experience` text DEFAULT NULL,
  `formation` text DEFAULT NULL,
  `path` varchar(100) DEFAULT NULL,
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
(21,'[\"{\\\"name\\\":\\\"grdgr\\\",\\\"description\\\":\\\"grd\\\",\\\"year\\\":\\\"grd\\\"}\"]','[\"{\\\"enterprise\\\":\\\"gdrgrd\\\",\\\"job\\\":\\\"grdgr\\\",\\\"year\\\":\\\"févr.-2022\\\\nau\\\\ndéc.-2022\\\",\\\"competences\\\":[\\\"grdgrdgrd\\\"]}\"]','[\"qsdfgdg\"]',NULL,NULL),
(22,'[\"{\\\"name\\\":\\\"zdq\\\",\\\"description\\\":\\\"dzq\\\",\\\"year\\\":\\\"dzq\\\"}\"]','[\"{\\\"enterprise\\\":\\\"dzq\\\",\\\"job\\\":\\\"dzq\\\",\\\"year\\\":\\\"févr.-2022\\\\nau\\\\ndéc.-2022\\\",\\\"competences\\\":[]}\"]','[\"fefesfes\"]',NULL,NULL),
(23,'[\"{\\\"name\\\":\\\"fes\\\",\\\"description\\\":\\\"fesfes\\\",\\\"year\\\":\\\"fesf\\\"}\",\"{\\\"name\\\":\\\"fesfes\\\",\\\"description\\\":\\\"fesfes\\\",\\\"year\\\":\\\"fesf\\\"}\",\"{\\\"name\\\":\\\"fes\\\",\\\"description\\\":\\\"fes\\\",\\\"year\\\":\\\"fes\\\"}\"]','[\"{\\\"enterprise\\\":\\\"fesfes\\\",\\\"job\\\":\\\"fesefs\\\",\\\"year\\\":\\\"janv.-2022\\\\nau\\\\njanv.-2022\\\",\\\"competences\\\":[\\\"fesfes\\\"]}\",\"{\\\"enterprise\\\":\\\"fesfes\\\",\\\"job\\\":\\\"fesfesfes\\\",\\\"year\\\":\\\"juin-2022\\\\nau\\\\nfévr.-2022\\\",\\\"competences\\\":[\\\"fessf\\\"]}\"]','[\"fsefse\"]',NULL,NULL);
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
  `event_type` enum('presence','absence','prestation suplémentaire','astreinte','autre') DEFAULT NULL,
  `comment` text DEFAULT NULL,
  `other_event` enum('IM','FI','VM') DEFAULT NULL,
  `value` double NOT NULL,
  `consultant_backup` varchar(100) DEFAULT NULL,
  `absence_event` enum('Congés payé','Congés sans solde','RTT','Maladie','Récupération','Evenement famillial','Enfant malade','Paternité') DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `Calendar_FK` (`user_id`),
  CONSTRAINT `Calendar_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=416 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Calendar`
--

LOCK TABLES `Calendar` WRITE;
/*!40000 ALTER TABLE `Calendar` DISABLE KEYS */;
INSERT INTO `Calendar` VALUES
(7,'2022-11-01','presence','',NULL,1,'',NULL,394),
(7,'2022-11-02','presence','',NULL,1,'',NULL,395),
(7,'2022-11-03','presence','',NULL,1,'',NULL,396),
(7,'2022-11-04','presence','',NULL,1,'',NULL,397),
(7,'2022-11-07','presence','',NULL,1,'',NULL,398),
(7,'2022-11-08','presence','',NULL,1,'',NULL,399),
(7,'2022-11-09','presence','',NULL,1,'',NULL,400),
(7,'2022-11-10','presence','',NULL,1,'',NULL,401),
(7,'2022-11-11','presence','',NULL,1,'',NULL,402),
(7,'2022-11-14','presence','',NULL,1,'',NULL,403),
(7,'2022-11-15','presence','',NULL,1,'',NULL,404),
(7,'2022-11-16','presence','',NULL,1,'',NULL,405),
(7,'2022-11-17','presence','',NULL,1,'',NULL,406),
(7,'2022-11-18','presence','',NULL,1,'',NULL,407),
(7,'2022-11-21','presence','',NULL,1,'',NULL,408),
(7,'2022-11-22','presence','',NULL,1,'',NULL,409),
(7,'2022-11-23','presence','',NULL,1,'',NULL,410),
(7,'2022-11-24','presence','',NULL,1,'',NULL,411),
(7,'2022-11-25','presence','',NULL,1,'',NULL,412),
(7,'2022-11-28','presence','',NULL,1,'',NULL,413),
(7,'2022-11-29','presence','',NULL,1,'',NULL,414),
(7,'2022-11-30','presence','',NULL,1,'',NULL,415);
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
  `firstname` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` varchar(12) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Candidat`
--

LOCK TABLES `Candidat` WRITE;
/*!40000 ALTER TABLE `Candidat` DISABLE KEYS */;
INSERT INTO `Candidat` VALUES
(21,'fesfse','defs','remy@gmail.fr','fesfes'),
(22,'fse','fes','remysalem@hotmail.fr','fesfes'),
(23,'dzqdzqdzq','dzqdzqdzqdzqdzq','dzqdzqdzq@gmail.fr','dzqdqz');
/*!40000 ALTER TABLE `Candidat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Holliday`
--

DROP TABLE IF EXISTS `Holliday`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Holliday` (
  `user_id` int(11) DEFAULT NULL,
  `dateStart` date NOT NULL,
  `status` enum('pending','accepted','refused') DEFAULT 'pending',
  `dateEnd` date NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `Vacation_FK` (`user_id`),
  CONSTRAINT `Vacation_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Holliday`
--

LOCK TABLES `Holliday` WRITE;
/*!40000 ALTER TABLE `Holliday` DISABLE KEYS */;
INSERT INTO `Holliday` VALUES
(7,'2022-11-04','accepted','2022-11-18',35),
(7,'2022-11-10','refused','2022-11-18',57);
/*!40000 ALTER TABLE `Holliday` ENABLE KEYS */;
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
  `date` datetime DEFAULT NULL,
  `appreciation` text DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `RDV_FK` (`user_id`),
  KEY `RDV_FK_1` (`candidat_id`),
  CONSTRAINT `RDV_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`),
  CONSTRAINT `RDV_FK_1` FOREIGN KEY (`candidat_id`) REFERENCES `Candidat` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
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
  `uuid` varchar(100) DEFAULT NULL,
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
(7,'2022-11-25','39eff1e6-8238-4b30-8543-28dcb6081149'),
(7,'2022-11-30','9a96da9c-22f3-4453-bfcd-4cc4fc00745f');
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
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `role` enum('user','owner','admin') NOT NULL,
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES
(7,'Remy','098f6bcd4621d373cade4e832627b4f6','admin','remysalem@hotmail.fr'),
(36,'test','f19ff23b6036fabf08f7caad86ad085d','user','test@gmail.com'),
(37,'remy2','f19ff23b6036fabf08f7caad86ad085d','user','remy2@gmail.com');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-11-28 16:48:16
