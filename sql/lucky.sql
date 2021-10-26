/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : lucky

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 14/10/2021 08:26:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for desire
-- ----------------------------
DROP TABLE IF EXISTS `desire`;
CREATE TABLE `desire` (
  `id` int NOT NULL AUTO_INCREMENT,
  `desire` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `state` int NOT NULL,
  `type` int NOT NULL,
  `wishman_qq` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `wishman_wechat` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `wishman_tel` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `wishman_name` varchar(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `creat_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `light_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `light_user` int DEFAULT '-1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int NOT NULL AUTO_INCREMENT,
  `desire_id` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `student_number` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `email` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `tel` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `wechat` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `school` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci NOT NULL,
  `name` char(50) COLLATE utf8mb4_german2_ci DEFAULT NULL,
  `gender` tinyint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

-- ----------------------------
-- Table structure for user_desire
-- ----------------------------
DROP TABLE IF EXISTS `user_desire`;
CREATE TABLE `user_desire` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `desire_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_german2_ci;

SET FOREIGN_KEY_CHECKS = 1;
