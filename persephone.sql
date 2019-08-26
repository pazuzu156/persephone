-- MySQL dump 10.16  Distrib 10.1.40-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: persephone
-- ------------------------------------------------------
-- Server version	10.1.40-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `artist_image`
--

DROP TABLE IF EXISTS `artist_image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `artist_image` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `artist` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ma_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `artist_image`
--

LOCK TABLES `artist_image` WRITE;
/*!40000 ALTER TABLE `artist_image` DISABLE KEYS */;
/*!40000 ALTER TABLE `artist_image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `crown`
--

DROP TABLE IF EXISTS `crown`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `crown` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `discord_id` bigint(20) NOT NULL,
  `artist` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `play_count` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `crown`
--

LOCK TABLES `crown` WRITE;
/*!40000 ALTER TABLE `crown` DISABLE KEYS */;
INSERT INTO `crown` VALUES (2,192548928804225025,'Kataklysm',83),(3,192548928804225025,'Kampfar',59),(4,335441439238389761,'Darkthrone',570),(5,192548928804225025,'God Dethroned',83),(6,192548928804225025,'Armored Dawn',78),(7,192548928804225025,'Istapp',30),(8,192548928804225025,'Behemoth',31),(9,192548928804225025,'Venom',73),(10,192548928804225025,'Zornheym',44),(11,192548928804225025,'Equilibrium',25),(12,303528663008411650,'Exmortus',119),(13,192548928804225025,'Suicidal Angels',101),(14,192548928804225025,'HATH',30),(15,192548928804225025,'Hate',48),(16,366645916339535872,'Morbid Angel',109),(17,335441439238389761,'Cradle of Filth',243),(18,192548928804225025,'Sinsaenum',41),(19,192548928804225025,'Dark Oath',27);
/*!40000 ALTER TABLE `crown` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `discord_id` bigint(20) NOT NULL,
  `lastfm` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (2,'Sir_Calango',303528663008411650,'BasicMetalBitch'),(3,'nanner',157175128197824512,'abbathismydad'),(4,'Risa',366645916339535872,'rsltm'),(6,'Apollyon',192548928804225025,'Pazuzu156'),(7,'Ghostly Espurr',335441439238389761,'GhostlyEspurr');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-08-26 17:44:06
