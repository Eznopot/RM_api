-- MariaDB dump 10.19  Distrib 10.9.3-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: db_RMS
-- ------------------------------------------------------
-- Server version	10.9.3-MariaDB

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
(10,'[{\"insert\":\"Developper flutter\\n\"}]','[{\"insert\":\"Abeeway\\nStiilt\\n\"}]','[{\"insert\":\"Epitech\\n\"}]',NULL,NULL),
(12,'[{\"insert\":\"· La maintenance, les Systèmes d\'exploitation, la virtualisation, ticketing GLPI, le langage informatique, Réseaux : Intervenir sur une infrastructure réseau [TCP/IP, NAT, Routage simple...],\\n· Serveurs : Mise en place et configuration de serveurs Windows, 2016/2019\\n· Datacenter, GHOST, Backup, Anti-virus DNS, DHCP, WDS, WSUS, AD (Comptes, GPO, Scripts)\\n· Gestion de parcs sous GLPI/OCS,\\n· Messagerie Outlook,\\n· Thunderbird,\\n· notions Exchange Compétences associées\\n· ITIL\\n\"}]','[{\"insert\":\"Déc 2021 – mai 2022 Total Energies – TCS France\\n\\nTechnicienne Helpdesk\\n\\n· Prise en main à distance,\\n\\n· déploiement,\\n\\n· rédaction de procédures,\\n\\n· ticketing SNow,\\n\\n· MS Office 365,\\n\\n· support aux utilisateurs Total Energies\\n\\nSept 2020 – Août 2021 DSI Sorbonnes Université\\n\\nTechnicienne Helpdesk\\n\\n· Prise en main à distance,\\n\\n· déploiement,\\n\\n· déplacement sur site,\\n\\n· rédaction de procédures,\\n\\n· ticketing GLPI, MS Office 16-19,\\n\\n· installation et configuration des postes sur Windows et Mac...\\n\\nJanv 2019 – Mai 2019 Distritec Emerainville\\n\\nAssistante facturation\\n\\nMars 2018 – Déc 2018 Liver France\\n\\nTechnicienne SAV\\n\\n· Gestion du SAV de tablettes vendues avec l’adhésion à des magazines\\n\\n· Rédaction de procédure\\n\\n· Prise en charge des utilisateurs\\n\\nOctoMai 2015 – Déc 2017 Belles fringues\\n\\nVendeuse\\n\\nDéc 2014 – Janv 2015 TGI de Meaux\\n\\nAgent administratif\\n\\nSept – Déc 2009 VANXY\\n\\nSecrétaire polyvalente\\n\\nMars 2007 – Avril 2008 Collège des 4 arpents – Lagny sur Marne\\n\\nEmployée vie scolaire\\n\\nMai 2000 – Nov 2004 Centre E.LECLERC\\n\\nHôtesse de caisse\\n\\nFORMA\\n\"}]','[{\"insert\":\"2019 AFPA Champs-Sur-Marne Formation qualifiante AFPA Champs-sur-Marne Durée :8 mois (1190 heures).\\n\\nNiveau de la Formation Validation visée : Titre professionnel de niveau 4 (Bac technique) de technicien/ne d\'assistance en informatique.\\n\\n2018 AFPA Champs-sur-Marne stage bureautique en E-earning Secrétariat assistanat PACK OFFICE ACCESS EXCEL WORD OUTLOOK WINDOWS 7 POWER POINT\\n\\n2014 GRETA Champs-sur-Marne\\n\\nFormation CIEL/SAGE comptabilité.\\n\\n2000 LPR SERMENAZ RILLIEUX LA PAPE Bac pro métiers de la comptabilité\\n\"}]',NULL,NULL),
(14,'[{\"insert\":\"Belle\\n\"}]','[{\"insert\":\"est aller a dublin\\n\"}]','[{\"insert\":\"EDJ de merde\\n\"}]',NULL,NULL);
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
  `event_type` enum('presence','absence','backup','prestation suplémentaire','astreinte','autre') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `comment` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `other_event` enum('IM','FI','VM') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` double NOT NULL,
  `consultant_backup` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `absence_event` enum('Congés payé','Congés sans solde','RTT','Maladie','Récupération','Evenement famillial','Enfant malade','Paternité') COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `Calendar_FK` (`user_id`),
  CONSTRAINT `Calendar_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=236 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Calendar`
--

LOCK TABLES `Calendar` WRITE;
/*!40000 ALTER TABLE `Calendar` DISABLE KEYS */;
INSERT INTO `Calendar` VALUES
(7,'2022-10-03','presence','',NULL,1,'',NULL,215),
(7,'2022-10-04','presence','',NULL,1,'',NULL,216),
(7,'2022-10-05','presence','',NULL,1,'',NULL,217),
(7,'2022-10-06','presence','',NULL,1,'',NULL,218),
(7,'2022-10-07','presence','',NULL,1,'',NULL,219),
(7,'2022-10-10','presence','',NULL,1,'',NULL,220),
(7,'2022-10-11','presence','',NULL,1,'',NULL,221),
(7,'2022-10-12','presence','',NULL,1,'',NULL,222),
(7,'2022-10-13','presence','',NULL,1,'',NULL,223),
(7,'2022-10-14','presence','',NULL,1,'',NULL,224),
(7,'2022-10-17','presence','',NULL,1,'',NULL,225),
(7,'2022-10-18','presence','',NULL,1,'',NULL,226),
(7,'2022-10-19','presence','',NULL,1,'',NULL,227),
(7,'2022-10-20','presence','',NULL,1,'',NULL,228),
(7,'2022-10-21','presence','',NULL,1,'',NULL,229),
(7,'2022-10-24','presence','',NULL,1,'',NULL,230),
(7,'2022-10-25','presence','',NULL,1,'',NULL,231),
(7,'2022-10-26','presence','',NULL,1,'',NULL,232),
(7,'2022-10-27','presence','',NULL,1,'',NULL,233),
(7,'2022-10-28','presence','',NULL,1,'',NULL,234),
(7,'2022-10-31','presence','',NULL,1,'',NULL,235);
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Candidat`
--

LOCK TABLES `Candidat` WRITE;
/*!40000 ALTER TABLE `Candidat` DISABLE KEYS */;
INSERT INTO `Candidat` VALUES
(10,'Remy','Salem','remy@gmail'),
(12,'Serge','Lebg','serge@gmail.fr'),
(14,'Floflo','Amenta','amenta@gmail.com');
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
  `status` enum('pending','accepted','refused') COLLATE utf8mb4_unicode_ci DEFAULT 'pending',
  `dateEnd` date NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `Vacation_FK` (`user_id`),
  CONSTRAINT `Vacation_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Holliday`
--

LOCK TABLES `Holliday` WRITE;
/*!40000 ALTER TABLE `Holliday` DISABLE KEYS */;
INSERT INTO `Holliday` VALUES
(7,'2022-11-04','accepted','2022-11-18',35);
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
  `appreciation` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `RDV_FK` (`user_id`),
  KEY `RDV_FK_1` (`candidat_id`),
  CONSTRAINT `RDV_FK` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`),
  CONSTRAINT `RDV_FK_1` FOREIGN KEY (`candidat_id`) REFERENCES `Candidat` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `RDV`
--

LOCK TABLES `RDV` WRITE;
/*!40000 ALTER TABLE `RDV` DISABLE KEYS */;
INSERT INTO `RDV` VALUES
(7,14,'2022-11-14 00:30:00',NULL,3);
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
(7,'2022-11-09','57c5e966-e65e-45cb-8cf5-9705b201cd1a'),
(7,'2022-11-09','4e47fbea-58d0-4260-bdaf-7031bfe8c35f'),
(7,'2022-11-09','ffc9c0c8-370f-405d-85b2-f59248d358c4'),
(7,'2022-11-09','266a9989-a34b-4da6-a37c-f5c6f4645fdc'),
(7,'2022-11-09','b821b9f8-5f06-4553-b2e2-929bb1ebef14'),
(7,'2022-11-09','2b3bd961-0bb9-47ef-828d-dee17c4db58c'),
(7,'2022-11-09','50688c72-8237-433d-b73d-08e69162f6c8'),
(7,'2022-11-09','8c9e879f-c7e4-4be3-a4fb-f1b91d746ad9'),
(7,'2022-11-09','d681816d-cb20-4808-80d3-1e24545fa80c'),
(7,'2022-11-09','ef58b94a-8340-4e19-bc5f-38ae5423f919'),
(7,'2022-11-09','4098768b-06b8-4ea8-8bb1-1658a3375756'),
(7,'2022-11-09','aef9dab9-ed6f-423c-b2d5-0122b51093ca'),
(7,'2022-11-09','1044444a-d93f-4ec7-8ba3-d299fd08dd47'),
(7,'2022-11-09','43c82b31-5296-401c-b2fd-aa2558888e04'),
(7,'2022-11-09','da0c2f10-b880-45df-9264-3003c82acac5'),
(7,'2022-11-09','fd68881c-1d1d-4518-a859-e01889ab4fc3'),
(7,'2022-11-09','48620072-c69c-47cc-b68f-b66e7057e89d'),
(7,'2022-11-09','845d5878-1af9-47dd-b5e3-3ec8dff137ca'),
(7,'2022-11-09','f5050a76-63f2-42db-aa85-cbf1d9a125ee'),
(7,'2022-11-09','cdcb98ed-e6cc-4662-ac24-0e21fcddb2d9'),
(7,'2022-11-09','335b2c97-c55d-4ed6-b6e7-b58b00a86d94'),
(7,'2022-11-09','28109172-55e2-43d9-862a-38f8c016a533'),
(7,'2022-11-09','0a5dc1f2-d92b-4bc0-8c83-f645ee472901'),
(7,'2022-11-09','28ea918f-2499-4d8d-9e7b-bd0c8e5d893d'),
(7,'2022-11-09','bb3d7187-8ed5-495b-b9cd-234148a98f12'),
(7,'2022-11-09','68184338-98d4-440d-8948-4e705153d389'),
(7,'2022-11-09','e4f2098c-56f7-4866-9b55-c610280f2827'),
(7,'2022-11-09','5898797f-078e-437c-b966-d481f5cb25be'),
(7,'2022-11-09','3ba42cbd-c618-4824-8e8e-315c6e45439c'),
(7,'2022-11-09','a0a06a78-827c-4dbb-b5a6-65aa2160418b'),
(7,'2022-11-09','375f4754-c54a-4379-874f-d1a662ad7fff'),
(7,'2022-11-09','c1cb4c95-c5ec-4d0c-90b7-9a5ea65282d3'),
(7,'2022-11-09','c5bd3335-6795-4e50-aca3-ac2454c3e9ab'),
(7,'2022-11-09','9941d1b3-85d6-4c61-a216-4e50327a39fa'),
(7,'2022-11-09','57021c16-edda-44b7-987d-ec9525e30237'),
(7,'2022-11-09','f23cfacb-97cf-4b14-be36-c89689cd9f75'),
(7,'2022-11-09','1e99726a-7886-4f44-86a6-aed257a35593'),
(7,'2022-11-09','97e7804b-b739-49e4-849d-d26936bcee99'),
(7,'2022-11-09','09a05b9f-b9ef-463d-8471-3444cb833c79'),
(7,'2022-11-09','195e6a00-17d1-4c15-8161-1ff288cfdc9c'),
(7,'2022-11-09','fba98a64-be1a-4121-bea3-969cb09cbbd9'),
(7,'2022-11-09','6edf4e59-63e0-43c0-b593-a7bfd2a52a17'),
(7,'2022-11-09','c774694c-6d59-4e20-853d-7ff58a06ec82'),
(7,'2022-11-09','9322535f-e643-483d-bc12-495dfb6a1b7d'),
(7,'2022-11-09','9ad29264-bac2-4d16-b128-77b35e1f3482'),
(7,'2022-11-09','1624ff10-3f60-4cac-a8d7-95c0db2a4887'),
(7,'2022-11-09','d84c2103-0edb-42d8-8a89-896571ae999e'),
(7,'2022-11-09','9300266b-669d-406d-ba45-b087b01a6279'),
(7,'2022-11-09','9a41eeb9-980e-444c-a084-34e35c55dd98'),
(7,'2022-11-09','0a8fd30c-ef23-4589-bec9-df7d3dd9365b'),
(7,'2022-11-09','b417f463-c9af-406e-ba6f-e1c2a2e7fcb6'),
(7,'2022-11-09','46e3fbb5-74d2-4241-b855-a98518a831a4'),
(7,'2022-11-09','21d29e6b-775c-4f6b-98d3-2916ff4ae504'),
(7,'2022-11-09','62b652f7-ff8a-4424-9eca-1376a0535cf9'),
(7,'2022-11-09','045a2418-cb21-4a72-90d7-e025d8d44d40'),
(7,'2022-11-09','f8a3d37c-3ee7-427f-bca6-75daa6965f34'),
(7,'2022-11-09','e34e48b7-17fb-4a73-a3a6-17c69e2b644e'),
(7,'2022-11-09','87f8ce0b-1942-4ddb-8f61-dc79c63c3f8e'),
(7,'2022-11-09','c8a0ed79-e581-4a99-aeff-56d3a0611963'),
(7,'2022-11-09','a12416fc-fec2-4adf-8c0c-d69738c8f0f2'),
(7,'2022-11-09','cb0c4074-393e-4c0b-9bfd-5b9bedb5faf0'),
(7,'2022-11-09','4998f9bc-a71d-4d47-b143-6dfee1b2962f'),
(7,'2022-11-09','942a6ec5-fc55-4383-85b1-048cbc45b57e'),
(7,'2022-11-09','afe17e7b-e2f7-4ba6-990c-df06796f3baa'),
(7,'2022-11-09','78d495f2-b9fb-4183-a70a-9090ed28d918'),
(7,'2022-11-09','9d65718a-4967-4623-b99f-70175f25f611'),
(7,'2022-11-09','7680d914-0dab-414a-9403-92a42bd6cc22'),
(7,'2022-11-09','477da762-f100-4c26-88f4-ab2d49def29e'),
(7,'2022-11-09','2e136b27-a789-4fac-ae2c-acef29ce171a'),
(7,'2022-11-09','25c48ff9-fe44-4801-a273-79868f33b0c0'),
(7,'2022-11-09','0b3b69fa-9912-4844-bff9-849a58a4cbe0'),
(7,'2022-11-09','60942f6b-63ab-4082-9d40-2a38f61a9290'),
(7,'2022-11-09','2c8ffb67-33ff-49b1-8808-ac763699d0bc'),
(7,'2022-11-09','5b432667-9d8f-4752-bdd7-ac6ce21511cb'),
(7,'2022-11-09','bd67f9fa-a4e8-45cd-8da9-a1a5792ee54a'),
(7,'2022-11-09','ae958dc6-e45e-4325-93c8-4ab03de5a60e'),
(7,'2022-11-09','37bda10c-d311-4d8d-ab90-65dfbf3b5c76'),
(7,'2022-11-09','77a98cc2-7993-4235-aa69-791381838816'),
(7,'2022-11-09','4561f824-3341-476f-bc99-07c672d84581'),
(7,'2022-11-09','cb74ccd9-e2b0-4694-bdd8-42b072ed0b55'),
(7,'2022-11-09','d55bbbea-a8da-4398-89a3-8b3dc6bd0b11'),
(7,'2022-11-09','7eaacbee-e216-4ef0-bafb-f9a81f2aff04'),
(7,'2022-11-09','3498b889-4e8e-47ca-9973-b87ab0d76eae'),
(7,'2022-11-09','0d52a042-e147-40e5-b11e-fe418c62fdd7'),
(7,'2022-11-09','1726ec48-7650-4d54-a055-6446973ad188'),
(7,'2022-11-09','3de1f05a-b4f6-4fb5-a983-111d500e96d9'),
(7,'2022-11-09','8600cc49-5468-4616-9fb2-d862e927a6d9'),
(7,'2022-11-09','a2f7220c-4c90-4419-8d14-ba986d7604cb'),
(7,'2022-11-09','29832022-ed4b-4d42-8ce7-491b258f89b9'),
(7,'2022-11-09','00d6cb7c-726b-42e0-b743-f1ac02286ef4'),
(7,'2022-11-09','fe3e8838-797d-4d93-8a16-e974a31027c5'),
(7,'2022-11-09','d802ba05-6993-46ec-8487-e17399b33ab8'),
(7,'2022-11-09','d88c84ca-d07d-446c-8430-46988decc14a'),
(7,'2022-11-09','d32d91db-3c65-4032-9e16-f1fb2e33c1bc'),
(7,'2022-11-09','a2ce6744-cb4f-465c-9a9d-e8568ca94c8b'),
(7,'2022-11-09','cad9226c-dbbb-4d02-9e07-8ee4ffd6c5a1'),
(7,'2022-11-09','912071a7-7f48-4d22-8c1f-7e3e493ba29b'),
(7,'2022-11-09','e720035e-f7ca-4dd0-8772-8c905fd9e678'),
(7,'2022-11-09','07fad116-989c-49c8-ad3e-5f80df6fb6e0'),
(7,'2022-11-09','1ebc32ef-88d3-4113-b38d-00116d397682'),
(7,'2022-11-09','89655d10-2349-4503-9603-2115c938f0df'),
(7,'2022-11-09','ab8b0c56-562e-4624-be88-1c46cfaaa6d3'),
(7,'2022-11-09','d1b8e93b-ad4c-463e-b7e3-823e19bdb008'),
(7,'2022-11-09','87be3648-afe7-41ef-b78a-ebebf80c574f'),
(7,'2022-11-09','6dda20ab-9242-4cbc-8be3-195c92ef4ad4'),
(7,'2022-11-09','c1434fe1-4655-475a-9bc4-52ce42cffb49'),
(7,'2022-11-09','147f400f-8149-4c3e-9639-72685de05161'),
(7,'2022-11-09','0096a15c-afe7-48aa-9f0a-7a40a7c6967d'),
(7,'2022-11-09','6d686d45-fdd6-4386-805f-618c2e0edf71'),
(7,'2022-11-09','a093f94b-6956-40af-a573-d5328be5a520'),
(7,'2022-11-09','1b797c45-3132-49c3-b5f6-6e79ba4d2b53'),
(7,'2022-11-09','84cef11b-af5b-4061-9ef7-454595934201'),
(7,'2022-11-16','2a7f9620-c0de-4436-ae72-01cb31e68488');
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

-- Dump completed on 2022-11-14 16:46:25
