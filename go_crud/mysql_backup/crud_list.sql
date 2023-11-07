/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : crud-list

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 07/11/2023 22:47:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for crud_list
-- ----------------------------
DROP TABLE IF EXISTS `crud_list`;
CREATE TABLE `crud_list`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `birthday` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_crud_list_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of crud_list
-- ----------------------------
INSERT INTO `crud_list` VALUES (1, '2023-11-07 21:27:35.394', '2023-11-07 21:27:35.394', '2023-11-07 22:45:32.469', '张三', '', '', '', '', '');
INSERT INTO `crud_list` VALUES (2, '2023-11-07 21:31:27.696', '2023-11-07 21:31:27.696', NULL, '张三', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (3, '2023-11-07 21:31:32.675', '2023-11-07 21:31:32.675', NULL, '张三', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (4, '2023-11-07 21:37:03.731', '2023-11-07 21:37:03.731', '2023-11-07 22:46:06.996', '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (5, '2023-11-07 21:37:17.170', '2023-11-07 21:37:17.170', NULL, '李四', '55', '2222@22.mail', '', '', '');
INSERT INTO `crud_list` VALUES (6, '2023-11-07 21:39:26.496', '2023-11-07 21:39:26.496', NULL, '李四', '55', '2222@22.mail', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
