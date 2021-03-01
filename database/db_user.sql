-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: data_users
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `m_pekerjaan`
--

LOCK TABLES `m_pekerjaan` WRITE;
/*!40000 ALTER TABLE `m_pekerjaan` DISABLE KEYS */;
INSERT INTO `m_pekerjaan` VALUES ('8219c7cc-30bb-11eb-b405-c85b766bafe8','PNS'),('821e309f-30bb-11eb-b405-c85b766bafe8','Wirausaha'),('82225b07-30bb-11eb-b405-c85b766bafe8','Wiraswasta');
/*!40000 ALTER TABLE `m_pekerjaan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `m_pendidikan`
--

LOCK TABLES `m_pendidikan` WRITE;
/*!40000 ALTER TABLE `m_pendidikan` DISABLE KEYS */;
INSERT INTO `m_pendidikan` VALUES ('5ddbb91e-30bb-11eb-b405-c85b766bafe8','SD'),('5de17471-30bb-11eb-b405-c85b766bafe8','SMP'),('5de57c7c-30bb-11eb-b405-c85b766bafe8','SMA'),('5de9a553-30bb-11eb-b405-c85b766bafe8','SD'),('5dedbec5-30bb-11eb-b405-c85b766bafe8','Diploma'),('5df1331d-30bb-11eb-b405-c85b766bafe8','Sarjana'),('5df4cbfb-30bb-11eb-b405-c85b766bafe8','Magister'),('5df8d140-30bb-11eb-b405-c85b766bafe8','Doktor');
/*!40000 ALTER TABLE `m_pendidikan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `tb_users`
--

LOCK TABLES `tb_users` WRITE;
/*!40000 ALTER TABLE `tb_users` DISABLE KEYS */;
INSERT INTO `tb_users` VALUES ('36d0035f-cd53-4b55-a50c-0ce77d7b8ee4','99999999999','DIAN AFRILIAN NOERY','1995-04-29','82225b07-30bb-11eb-b405-c85b766bafe8','5df1331d-30bb-11eb-b405-c85b766bafe8',1,'2020-12-01','2020-12-01');
/*!40000 ALTER TABLE `tb_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-01  4:01:49

CREATE TABLE `data_users`.`user_login` (
  `iduser_login` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(75) NULL,
  `password` VARCHAR(80) NULL,
  PRIMARY KEY (`iduser_login`));
