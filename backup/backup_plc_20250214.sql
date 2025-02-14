/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-11.6.2-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: plc_config
-- ------------------------------------------------------
-- Server version	11.6.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `logs`
--

DROP TABLE IF EXISTS `logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `timestamp` datetime NOT NULL,
  `level` varchar(10) NOT NULL,
  `message` text NOT NULL,
  `additional` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=243 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `logs`
--

LOCK TABLES `logs` WRITE;
/*!40000 ALTER TABLE `logs` DISABLE KEYS */;
INSERT INTO `logs` VALUES
(1,'2025-02-14 15:57:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(2,'2025-02-14 15:57:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(3,'2025-02-14 15:58:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(4,'2025-02-14 15:58:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(5,'2025-02-14 15:58:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(6,'2025-02-14 15:59:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(7,'2025-02-14 15:59:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(8,'2025-02-14 15:59:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(9,'2025-02-14 16:00:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(10,'2025-02-14 16:00:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(11,'2025-02-14 16:00:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(12,'2025-02-14 16:01:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(13,'2025-02-14 16:01:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(14,'2025-02-14 16:01:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(15,'2025-02-14 16:02:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(16,'2025-02-14 16:02:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(17,'2025-02-14 16:02:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(18,'2025-02-14 16:03:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(19,'2025-02-14 16:03:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(20,'2025-02-14 16:03:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(21,'2025-02-14 16:04:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(22,'2025-02-14 16:04:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(23,'2025-02-14 16:04:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(24,'2025-02-14 16:05:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(25,'2025-02-14 16:05:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(26,'2025-02-14 16:05:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(27,'2025-02-14 16:06:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(28,'2025-02-14 16:06:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(29,'2025-02-14 16:06:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(30,'2025-02-14 16:07:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(31,'2025-02-14 16:07:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(32,'2025-02-14 16:07:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(33,'2025-02-14 16:08:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(34,'2025-02-14 16:08:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(35,'2025-02-14 16:08:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(36,'2025-02-14 16:09:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(37,'2025-02-14 16:09:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(38,'2025-02-14 16:09:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(39,'2025-02-14 16:10:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(40,'2025-02-14 16:10:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(41,'2025-02-14 16:10:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(42,'2025-02-14 16:11:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(43,'2025-02-14 16:11:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(44,'2025-02-14 16:11:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(45,'2025-02-14 16:12:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(46,'2025-02-14 16:12:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(47,'2025-02-14 16:12:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(48,'2025-02-14 16:13:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(49,'2025-02-14 16:13:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(50,'2025-02-14 16:13:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(51,'2025-02-14 16:14:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(52,'2025-02-14 16:14:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(53,'2025-02-14 16:14:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(54,'2025-02-14 16:15:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(55,'2025-02-14 16:15:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(56,'2025-02-14 16:15:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(57,'2025-02-14 16:16:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(58,'2025-02-14 16:16:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(59,'2025-02-14 16:16:41','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(60,'2025-02-14 16:17:01','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(61,'2025-02-14 16:17:13','INFO','Shutdown','Sinal de shutdown recebido'),
(62,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG4'),
(63,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG8'),
(64,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG2'),
(65,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG3'),
(66,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG7'),
(67,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG5'),
(68,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG6'),
(69,'2025-02-14 16:17:13','WARN','Encerramento do monitoramento da tag','TAG1'),
(70,'2025-02-14 16:17:15','WARN','Shutdown do gerenciamento de tags','PLC ID: 1'),
(71,'2025-02-14 16:17:15','WARN','Conexão perdida','PLC1'),
(72,'2025-02-14 16:17:21','ERROR','Erro ao conectar ao PLC','PLC2: dial tcp 192.168.0.102:102: i/o timeout'),
(73,'2025-02-14 16:17:31','INFO','Aplicação encerrada',''),
(74,'2025-02-14 16:17:57','INFO','Conectado ao PLC','PLC1'),
(75,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG7'),
(76,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG8'),
(77,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG1'),
(78,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG2'),
(79,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG3'),
(80,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG4'),
(81,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG5'),
(82,'2025-02-14 16:17:57','INFO','Iniciou monitoramento da tag','TAG6'),
(83,'2025-02-14 16:18:27','INFO','Status do PLC atualizado','PLC ID 1: online'),
(84,'2025-02-14 16:25:37','INFO','CreateTag','Tag criada: TAG10 para PLC 1'),
(85,'2025-02-14 16:25:42','INFO','Iniciou monitoramento da tag','TAG10'),
(86,'2025-02-14 16:25:44','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(87,'2025-02-14 16:25:46','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(88,'2025-02-14 16:25:48','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(89,'2025-02-14 16:25:50','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(90,'2025-02-14 16:25:52','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(91,'2025-02-14 16:25:54','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(92,'2025-02-14 16:25:56','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(93,'2025-02-14 16:25:58','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(94,'2025-02-14 16:26:00','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(95,'2025-02-14 16:26:02','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(96,'2025-02-14 16:26:04','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(97,'2025-02-14 16:26:06','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(98,'2025-02-14 16:26:08','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(99,'2025-02-14 16:26:10','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(100,'2025-02-14 16:26:12','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(101,'2025-02-14 16:26:14','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(102,'2025-02-14 16:26:16','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(103,'2025-02-14 16:26:18','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(104,'2025-02-14 16:26:20','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(105,'2025-02-14 16:26:22','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(106,'2025-02-14 16:26:24','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(107,'2025-02-14 16:26:26','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(108,'2025-02-14 16:26:28','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(109,'2025-02-14 16:26:30','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(110,'2025-02-14 16:26:32','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(111,'2025-02-14 16:26:34','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(112,'2025-02-14 16:26:36','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(113,'2025-02-14 16:26:36','INFO','UpdateTag','Tag atualizada: ID 14'),
(114,'2025-02-14 16:26:38','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(115,'2025-02-14 16:26:40','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(116,'2025-02-14 16:26:42','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(117,'2025-02-14 16:26:44','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(118,'2025-02-14 16:26:46','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(119,'2025-02-14 16:26:48','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(120,'2025-02-14 16:26:50','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(121,'2025-02-14 16:26:52','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(122,'2025-02-14 16:26:54','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(123,'2025-02-14 16:26:56','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(124,'2025-02-14 16:26:58','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(125,'2025-02-14 16:27:00','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(126,'2025-02-14 16:27:02','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(127,'2025-02-14 16:27:04','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(128,'2025-02-14 16:27:06','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(129,'2025-02-14 16:27:08','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(130,'2025-02-14 16:27:10','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(131,'2025-02-14 16:27:12','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(132,'2025-02-14 16:27:14','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(133,'2025-02-14 16:27:16','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(134,'2025-02-14 16:27:18','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(135,'2025-02-14 16:27:20','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(136,'2025-02-14 16:27:22','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(137,'2025-02-14 16:27:24','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(138,'2025-02-14 16:27:26','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(139,'2025-02-14 16:27:28','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(140,'2025-02-14 16:27:30','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(141,'2025-02-14 16:27:32','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(142,'2025-02-14 16:27:34','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(143,'2025-02-14 16:27:36','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(144,'2025-02-14 16:27:38','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(145,'2025-02-14 16:27:40','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(146,'2025-02-14 16:27:42','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(147,'2025-02-14 16:27:44','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(148,'2025-02-14 16:27:46','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(149,'2025-02-14 16:27:48','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(150,'2025-02-14 16:27:50','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(151,'2025-02-14 16:27:52','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(152,'2025-02-14 16:27:54','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(153,'2025-02-14 16:27:56','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(154,'2025-02-14 16:27:58','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(155,'2025-02-14 16:28:00','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(156,'2025-02-14 16:28:02','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(157,'2025-02-14 16:28:04','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(158,'2025-02-14 16:28:06','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(159,'2025-02-14 16:28:08','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(160,'2025-02-14 16:28:10','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(161,'2025-02-14 16:28:12','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(162,'2025-02-14 16:28:14','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(163,'2025-02-14 16:28:16','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(164,'2025-02-14 16:28:18','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(165,'2025-02-14 16:28:20','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(166,'2025-02-14 16:28:22','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(167,'2025-02-14 16:28:24','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(168,'2025-02-14 16:28:26','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(169,'2025-02-14 16:28:28','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(170,'2025-02-14 16:28:30','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(171,'2025-02-14 16:28:30','INFO','UpdateTag','Tag atualizada: ID 14'),
(172,'2025-02-14 16:28:32','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(173,'2025-02-14 16:28:34','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(174,'2025-02-14 16:28:36','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(175,'2025-02-14 16:28:38','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(176,'2025-02-14 16:28:40','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(177,'2025-02-14 16:28:42','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(178,'2025-02-14 16:28:44','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(179,'2025-02-14 16:28:46','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(180,'2025-02-14 16:28:48','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(181,'2025-02-14 16:28:50','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(182,'2025-02-14 16:28:52','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(183,'2025-02-14 16:28:54','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(184,'2025-02-14 16:28:56','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(185,'2025-02-14 16:28:58','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(186,'2025-02-14 16:29:00','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(187,'2025-02-14 16:29:02','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(188,'2025-02-14 16:29:04','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(189,'2025-02-14 16:29:06','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(190,'2025-02-14 16:29:08','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(191,'2025-02-14 16:29:10','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(192,'2025-02-14 16:29:12','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(193,'2025-02-14 16:29:14','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(194,'2025-02-14 16:29:16','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(195,'2025-02-14 16:29:18','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(196,'2025-02-14 16:29:20','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(197,'2025-02-14 16:29:22','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(198,'2025-02-14 16:29:24','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(199,'2025-02-14 16:29:26','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(200,'2025-02-14 16:29:28','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(201,'2025-02-14 16:29:30','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(202,'2025-02-14 16:29:32','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(203,'2025-02-14 16:29:34','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(204,'2025-02-14 16:29:35','INFO','UpdateTag','Tag atualizada: ID 14'),
(205,'2025-02-14 16:29:36','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(206,'2025-02-14 16:29:38','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(207,'2025-02-14 16:29:40','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(208,'2025-02-14 16:29:42','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(209,'2025-02-14 16:29:44','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(210,'2025-02-14 16:29:46','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(211,'2025-02-14 16:29:48','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(212,'2025-02-14 16:29:50','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(213,'2025-02-14 16:29:52','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(214,'2025-02-14 16:29:54','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(215,'2025-02-14 16:29:56','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(216,'2025-02-14 16:29:58','ERROR','Erro ao ler tag','TAG10: CPU : Address out of range'),
(217,'2025-02-14 16:29:59','INFO','Shutdown','Sinal de shutdown recebido'),
(218,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG4'),
(219,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG5'),
(220,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG7'),
(221,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG2'),
(222,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG6'),
(223,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG3'),
(224,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG8'),
(225,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG1'),
(226,'2025-02-14 16:29:59','WARN','Encerramento do monitoramento da tag','TAG10'),
(227,'2025-02-14 16:30:02','WARN','Shutdown do gerenciamento de tags','PLC ID: 1'),
(228,'2025-02-14 16:30:02','WARN','Conexão perdida','PLC1'),
(229,'2025-02-14 16:30:07','INFO','Aplicação encerrada',''),
(230,'2025-02-14 16:30:12','INFO','Conectado ao PLC','PLC1'),
(231,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG8'),
(232,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG1'),
(233,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG3'),
(234,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG7'),
(235,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG6'),
(236,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','SENSOR'),
(237,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG2'),
(238,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG4'),
(239,'2025-02-14 16:30:12','INFO','Iniciou monitoramento da tag','TAG5'),
(240,'2025-02-14 16:30:42','INFO','Status do PLC atualizado','PLC ID 1: online'),
(241,'2025-02-14 16:31:17','INFO','CreateTag','Tag criada: NIVEL para PLC 1'),
(242,'2025-02-14 16:31:22','INFO','Iniciou monitoramento da tag','NIVEL');
/*!40000 ALTER TABLE `logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `plcs`
--

DROP TABLE IF EXISTS `plcs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `plcs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `ip_address` varchar(15) NOT NULL,
  `rack` int(11) NOT NULL,
  `slot` int(11) NOT NULL,
  `active` tinyint(1) DEFAULT 1,
  `status` varchar(20) DEFAULT 'offline',
  `last_update` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `plcs`
--

LOCK TABLES `plcs` WRITE;
/*!40000 ALTER TABLE `plcs` DISABLE KEYS */;
INSERT INTO `plcs` VALUES
(1,'PLC1','192.168.0.33',0,1,1,'online','2025-02-14 17:11:12');
/*!40000 ALTER TABLE `plcs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `plc_id` int(11) DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `db_number` int(11) NOT NULL,
  `byte_offset` int(11) NOT NULL,
  `data_type` varchar(20) NOT NULL,
  `can_write` tinyint(1) DEFAULT 0,
  `scan_rate` int(11) DEFAULT 1000,
  `monitor_changes` tinyint(1) DEFAULT 0,
  `active` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `plc_id` (`plc_id`),
  CONSTRAINT `tags_ibfk_1` FOREIGN KEY (`plc_id`) REFERENCES `plcs` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES
(1,1,'TAG1',4,0,'real',0,2000,0,1),
(2,1,'TAG2',4,4,'int',0,5000,0,1),
(3,1,'TAG3',4,6,'word',0,2000,0,1),
(4,1,'TAG4',4,8,'bool',0,2000,0,1),
(5,1,'TAG5',4,10,'string',0,1000,0,1),
(6,1,'TAG6',4,266,'real',0,6000,1,1),
(7,1,'TAG7',4,270,'real',0,1000,1,1),
(12,1,'TAG8',4,274,'int',0,2000,0,1),
(14,1,'SENSOR',4,276,'real',0,2000,0,1),
(15,1,'NIVEL',4,280,'real',0,1000,0,1);
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(50) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
(1,'admin','$2a$10$zbgNPkf7yZ28V9ITlPFiJukXNln8S3zWMDzXsRWEDPmVpJgAPoU1i','superadmin','2025-02-12 15:34:54','2025-02-12 15:34:54'),
(2,'novoOperador','$2a$10$U0BOsnykCbSdD1RojNBAeezD4mxEUakVGSkKzg.4nxCms3upmXbpG','operator','2025-02-12 15:51:57','2025-02-12 15:51:57');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2025-02-14 17:37:10
