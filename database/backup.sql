/*
 Navicat MySQL Dump SQL

 Source Server         : SmartContract
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : database-smart-contract-trieuvynet-9dc9.f.aivencloud.com:28360
 Source Schema         : smart_contract

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 16/05/2024 23:11:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for attendances
-- ----------------------------
DROP TABLE IF EXISTS `attendances`;
CREATE TABLE "attendances" (
  "id" int NOT NULL AUTO_INCREMENT,
  "user_id" char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  "check_in" char(50) DEFAULT NULL,
  "check_out" char(50) DEFAULT NULL,
  "location" varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'Vinaconex Building - 47 Dien Bien Phu, Da kao, Dictrict 1, Ho Chi Minh city',
  "over_time" float DEFAULT '0',
  "bonus" float DEFAULT '0',
  "salary" float DEFAULT '0',
  "total_hour" float DEFAULT '0',
  PRIMARY KEY ("id","user_id")
);

-- ----------------------------
-- Records of attendances
-- ----------------------------
BEGIN;
INSERT INTO `attendances` (`id`, `user_id`, `check_in`, `check_out`, `location`, `over_time`, `bonus`, `salary`, `total_hour`) VALUES (5, '45e520a0942b47da84c2534e19927ffe', '2024/05/16 - 18:10:02', '2024/05/16 - 18:15:31', 'Vinaconex Building - 47 Dien Bien Phu, Da kao, Dictrict 1, Ho Chi Minh city', 0, 0, 0, 0.0913889);
INSERT INTO `attendances` (`id`, `user_id`, `check_in`, `check_out`, `location`, `over_time`, `bonus`, `salary`, `total_hour`) VALUES (6, '45e520a0942b47da84c2534e19927ffe', '2024/05/16 - 18:19:49', '2024/05/16 - 18:19:55', 'Vinaconex Building - 47 Dien Bien Phu, Da kao, Dictrict 1, Ho Chi Minh city', 0, 0, 0, 0.00166667);
INSERT INTO `attendances` (`id`, `user_id`, `check_in`, `check_out`, `location`, `over_time`, `bonus`, `salary`, `total_hour`) VALUES (7, '9b80dfa2814deab848494791980d7039', '2024/05/16 - 18:49:42', '2024/05/16 - 18:49:59', 'Vinaconex Building - 47 Dien Bien Phu, Da kao, Dictrict 1, Ho Chi Minh city', 0, 0, 0, 0.00472222);
INSERT INTO `attendances` (`id`, `user_id`, `check_in`, `check_out`, `location`, `over_time`, `bonus`, `salary`, `total_hour`) VALUES (8, '0899523aafc8cc44b8c1004ef84c845d', '2024/05/16 - 19:29:24', '2024/05/16 - 19:36:57', 'Vinaconex Building - 47 Dien Bien Phu, Da kao, Dictrict 1, Ho Chi Minh city', 0, 0, 0, 0.125833);
INSERT INTO `attendances` (`id`, `user_id`, `check_in`, `check_out`, `location`, `over_time`, `bonus`, `salary`, `total_hour`) VALUES (9, '0899523aafc8cc44b8c1004ef84c845d', '2024/05/16 - 19:37:13', '2024/05/16 - 19:37:24', 'Vinaconex Building - 47 Dien Bien Phu, Da kao, Dictrict 1, Ho Chi Minh city', 0, 0, 0, 0.00305556);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE "users" (
  "id" char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  "email" varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  "password" varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  "fullname" varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  "position" varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  "department" varchar(45) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `email`, `password`, `fullname`, `position`, `department`) VALUES ('0899523aafc8cc44b8c1004ef84c845d', 'trieuvynet03.1909@gmail.com', '$2a$10$rAYiUhsBlNwc2ryckg9CWeAhCaRCm3svOBlZPNG0HsHU.3j8H4VVG', 'Triệu Vỷ 02', 'Dev', 'IT');
INSERT INTO `users` (`id`, `email`, `password`, `fullname`, `position`, `department`) VALUES ('45e520a0942b47da84c2534e19927ffe', 'trieuvynet.1909@gmail.com', '$2a$12$m17AAT16bT2ay9o42v3x4.cMv82OLLWEYXD.M7Rbmib6oi9VDgepS', 'Lê Triệu Vỷ', 'Developer', 'IT');
INSERT INTO `users` (`id`, `email`, `password`, `fullname`, `position`, `department`) VALUES ('9b80dfa2814deab848494791980d7039', 'trieuvynet02.1909@gmail.com', '$2a$10$UFbChTonP.YkJgH/v5uKBOcA9GtiN4fmwgg6.gVd8EyQ5adIOJoq6', 'Triệu Vỷ 02', 'Dev', 'IT');
INSERT INTO `users` (`id`, `email`, `password`, `fullname`, `position`, `department`) VALUES ('af23211972a2bd3c55ec6597f76feaf0', 'trieuvynet@gmail.com', '$2a$10$wTj1.OQiqMLrp4UxcZxGOeAioRPneTQvB5igMdXEaPeF3TaUUZqHe', 'Triệu Vỷ 02', 'Dev', 'IT');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
