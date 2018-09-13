/*
 Navicat Premium Data Transfer

 Source Server         : 交易所
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : 47.52.46.151:3306
 Source Schema         : solian

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 07/09/2018 11:19:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for exchange_turnover
-- ----------------------------
DROP TABLE IF EXISTS `exchange_turnover`;
CREATE TABLE `exchange_turnover`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `exchange_id` int(11) NULL DEFAULT NULL COMMENT '交易所id',
  `hours_turnover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '24小时成交额',
  `week_turnover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '7天成交额',
  `month_turnover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '30天成交额',
  `create_time` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7346 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
