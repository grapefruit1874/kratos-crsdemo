/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : user

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 06/09/2022 20:18:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `status` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'tom', '2022-09-05 05:07:16', '2022-09-05 05:07:16', '1');
INSERT INTO `users` VALUES (3, 'jimmy', '2022-09-05 05:07:16', '2022-09-05 05:07:16', '0');
INSERT INTO `users` VALUES (4, 'you', '2022-09-05 05:07:16', '2022-09-05 05:07:16', '1');
INSERT INTO `users` VALUES (6, 'ju', '2022-09-05 05:07:16', '2022-09-05 05:07:16', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
