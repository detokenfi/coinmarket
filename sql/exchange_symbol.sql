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

 Date: 07/09/2018 11:19:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for exchange_symbol
-- ----------------------------
DROP TABLE IF EXISTS `exchange_symbol`;
CREATE TABLE `exchange_symbol`  (
  `symbol_id` int(11) NOT NULL AUTO_INCREMENT,
  `exchange_id` int(11) NULL DEFAULT NULL,
  `symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '交易对',
  `base_coin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '每个交易所所对应的基础币',
  `coin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '币种名称',
  `create_time` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `symbol_price` varchar(225) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '价格',
  `symbol_vol` varchar(225) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '24小时的交易量',
  `symbol_turnover` varchar(225) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '成交额',
  `accounted` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '占比',
  `base_rate` varchar(225) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '基础汇率',
  `rate_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`symbol_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2546 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '这个表是交易所和币种的中间表\r\n专门做交易所和币种的对应关系。还有交易所的交易对的更新' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
