-- -------------------------------------------------------------
-- TablePlus 4.8.2(436)
--
-- https://tableplus.com/
--
-- Database: muskong
-- Generation Time: 2022-07-27 14:21:07.8560
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `rules` json DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

DROP TABLE IF EXISTS `admin_rules`;
CREATE TABLE `admin_rules` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int unsigned NOT NULL DEFAULT '0' COMMENT '上级菜单',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'menu' COMMENT '类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由路径',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `menu_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '菜单类型:tab=选项卡,link=链接,iframe=Iframe',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Url',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `keepalive` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'deny' COMMENT '缓存:deny=关闭,allow=开启',
  `extend` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'none' COMMENT '扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sequence` int NOT NULL DEFAULT '0' COMMENT '权重(排序)',
  `active` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态:deny=禁用,allow=启用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `sequence` (`sequence`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单和权限规则表';

DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `roles` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表';

DROP TABLE IF EXISTS `bank_prcptcds`;
CREATE TABLE `bank_prcptcds` (
  `id` int NOT NULL AUTO_INCREMENT,
  `bankname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `bank_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '银行编号',
  `prcptcd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '行号',
  `city_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '所属地区编号',
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_code` (`bank_code`,`city_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='银行行号表';

DROP TABLE IF EXISTS `banks`;
CREATE TABLE `banks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `abbr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `state` int NOT NULL DEFAULT '1',
  `bankname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `logo1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `procode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='银行';

DROP TABLE IF EXISTS `configs`;
CREATE TABLE `configs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='配置信息';

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `story_attitude_id` int NOT NULL,
  `story_id` int NOT NULL,
  `user_id` int NOT NULL,
  `state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '状态: 未支付unpaid, 已支付paid, 结束finish, 退款申请refund, 已退款refunded',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_state` (`state`),
  KEY `index_user` (`user_id`),
  KEY `index_attitude` (`story_attitude_id`),
  KEY `index_story` (`story_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='故事订单表';

DROP TABLE IF EXISTS `stories`;
CREATE TABLE `stories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '编号',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '内容',
  `state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'one' COMMENT '状态: 第几集one..., 剧终:end',
  `rate` tinyint NOT NULL DEFAULT '0' COMMENT '赎回费',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '单价',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='故事表';

DROP TABLE IF EXISTS `story_attitudes`;
CREATE TABLE `story_attitudes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int NOT NULL,
  `story_id` int NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '状态: 允许allow, 禁止deny',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_pid` (`pid`),
  KEY `index_story_id` (`story_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='故事态度表';

DROP TABLE IF EXISTS `story_progress`;
CREATE TABLE `story_progress` (
  `id` int NOT NULL AUTO_INCREMENT,
  `story_id` int NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '内容',
  `date` datetime NOT NULL COMMENT '发生的时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_story_id` (`story_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='故事进度表';

DROP TABLE IF EXISTS `user_accounts`;
CREATE TABLE `user_accounts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `before` decimal(15,2) DEFAULT NULL COMMENT '变更前金额',
  `change` decimal(15,2) DEFAULT NULL COMMENT '变更金额',
  `after` decimal(15,2) DEFAULT NULL COMMENT '变更后金额',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '变更说明',
  `table` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '变更对应的表',
  `table_id` bigint DEFAULT NULL COMMENT '变更对应表的ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员余额变动表';

DROP TABLE IF EXISTS `user_auths`;
CREATE TABLE `user_auths` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'personal, company',
  `state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'pass, fail',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'API认证, person本站, face人脸核身',
  `identity_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '身份证号',
  `unified_social` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '统一社会信用代码',
  `front_photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '正面照',
  `back_photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '背面照',
  `hand_photo` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '手持照片',
  `unified_social_photo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '营业执照照片',
  `company_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '企业名称',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员实名认证表';

DROP TABLE IF EXISTS `user_banks`;
CREATE TABLE `user_banks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `accounts` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `bankid` int NOT NULL DEFAULT '0' COMMENT '-1支付宝 -2微信',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员银行账户表';

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `wechat_openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `account_amount` decimal(15,2) DEFAULT NULL COMMENT '账户金额',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQ_USERNAME` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='会员表';

INSERT INTO `admin_roles` (`id`, `name`, `rules`, `description`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'admin', '[0]', '管理员权限', '2022-07-03 06:22:53', '2022-07-27 13:03:59', NULL);

INSERT INTO `admin_rules` (`id`, `pid`, `type`, `title`, `link`, `path`, `icon`, `menu_type`, `url`, `component`, `keepalive`, `extend`, `remark`, `sequence`, `active`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, 'menu', '控制台', 'dashboard/dashboard', 'dashboard', 'fa fa-dashboard', 'tab', '', '/src/views/backend/dashboard.vue', 'allow', 'none', 'remark_text', 999, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(2, 0, 'menu_dir', '权限管理', 'auth', 'auth', 'fa fa-group', NULL, '', '', 'deny', 'none', '', 100, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(3, 2, 'menu', '角色组管理', 'auth/group', 'auth/group', 'fa fa-group', 'tab', '', '/src/views/backend/auth/group/index.vue', 'allow', 'none', '', 99, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(4, 3, 'button', '查看', 'auth/group/index', '', '', NULL, '', '', 'deny', 'none', '', 99, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(5, 3, 'button', '添加', 'auth/group/add', '', '', NULL, '', '', 'deny', 'none', '', 99, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(6, 3, 'button', '编辑', 'auth/group/edit', '', '', NULL, '', '', 'deny', 'none', '', 99, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(7, 3, 'button', '删除', 'auth/group/del', '', '', NULL, '', '', 'deny', 'none', '', 99, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(8, 2, 'menu', '管理员管理', 'auth/admin', 'auth/admin', 'el-icon-UserFilled', 'tab', '', '/src/views/backend/auth/admin/index.vue', 'allow', 'none', '', 98, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(9, 8, 'button', '查看', 'auth/admin/index', '', '', NULL, '', '', 'deny', 'none', '', 98, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(10, 8, 'button', '添加', 'auth/admin/add', '', '', NULL, '', '', 'deny', 'none', '', 98, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(11, 8, 'button', '编辑', 'auth/admin/edit', '', '', NULL, '', '', 'deny', 'none', '', 98, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(12, 8, 'button', '删除', 'auth/admin/del', '', '', NULL, '', '', 'deny', 'none', '', 98, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(13, 2, 'menu', '菜单规则管理', 'auth/menu', 'auth/menu', 'el-icon-Grid', 'tab', '', '/src/views/backend/auth/menu/index.vue', 'allow', 'none', '', 97, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(14, 13, 'button', '查看', 'auth/menu/index', '', '', NULL, '', '', 'deny', 'none', '', 97, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(15, 13, 'button', '添加', 'auth/menu/add', '', '', NULL, '', '', 'deny', 'none', '', 97, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(16, 13, 'button', '编辑', 'auth/menu/edit', '', '', NULL, '', '', 'deny', 'none', '', 97, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(17, 13, 'button', '删除', 'auth/menu/del', '', '', NULL, '', '', 'deny', 'none', '', 97, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(18, 13, 'button', '快速排序', 'auth/menu/sortable', '', '', NULL, '', '', 'deny', 'none', '', 97, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(19, 2, 'menu', '管理员日志管理', 'auth/adminLog', 'auth/adminLog', 'el-icon-List', 'tab', '', '/src/views/backend/auth/adminLog/index.vue', 'allow', 'none', '', 96, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(20, 19, 'button', '查看', 'auth/adminLog/index', '', '', NULL, '', '', 'deny', 'none', '', 96, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(21, 0, 'menu_dir', '会员管理', 'user', 'user', 'fa fa-drivers-license', NULL, '', '', 'deny', 'none', '', 95, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(22, 21, 'menu', '会员管理', 'user/user', 'user/user', 'fa fa-user', 'tab', '', '/src/views/backend/user/user/index.vue', 'allow', 'none', '', 94, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(23, 22, 'button', '查看', 'user/user/index', '', '', NULL, '', '', 'deny', 'none', '', 94, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(24, 22, 'button', '添加', 'user/user/add', '', '', NULL, '', '', 'deny', 'none', '', 94, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(25, 22, 'button', '编辑', 'user/user/edit', '', '', NULL, '', '', 'deny', 'none', '', 94, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(26, 22, 'button', '删除', 'user/user/del', '', '', NULL, '', '', 'deny', 'none', '', 94, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(27, 21, 'menu', '会员分组管理', 'user/group', 'user/group', 'fa fa-group', 'tab', '', '/src/views/backend/user/group/index.vue', 'allow', 'none', '', 93, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(28, 27, 'button', '查看', 'user/group/index', '', '', NULL, '', '', 'deny', 'none', '', 93, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(29, 27, 'button', '添加', 'user/group/add', '', '', NULL, '', '', 'deny', 'none', '', 93, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(30, 27, 'button', '编辑', 'user/group/edit', '', '', NULL, '', '', 'deny', 'none', '', 93, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(31, 27, 'button', '删除', 'user/group/del', '', '', NULL, '', '', 'deny', 'none', '', 93, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(32, 21, 'menu', '会员规则管理', 'user/rule', 'user/rule', 'fa fa-th-list', 'tab', '', '/src/views/backend/user/rule/index.vue', 'allow', 'none', '', 92, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(33, 32, 'button', '查看', 'user/rule/index', '', '', NULL, '', '', 'deny', 'none', '', 92, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(34, 32, 'button', '添加', 'user/rule/add', '', '', NULL, '', '', 'deny', 'none', '', 92, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(35, 32, 'button', '编辑', 'user/rule/edit', '', '', NULL, '', '', 'deny', 'none', '', 92, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(36, 32, 'button', '删除', 'user/rule/del', '', '', NULL, '', '', 'deny', 'none', '', 92, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(37, 32, 'button', '快速排序', 'user/rule/sortable', '', '', NULL, '', '', 'deny', 'none', '', 92, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(38, 21, 'menu', '会员余额管理', 'user/moneyLog', 'user/moneyLog', 'el-icon-Money', 'tab', '', '/src/views/backend/user/moneyLog/index.vue', 'deny', 'none', '', 91, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(39, 38, 'button', '查看', 'user/moneyLog/index', '', '', NULL, '', '', 'deny', 'none', '', 91, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(40, 38, 'button', '添加', 'user/moneyLog/add', '', '', NULL, '', '', 'deny', 'none', '', 91, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(41, 21, 'menu', '会员积分管理', 'user/scoreLog', 'user/scoreLog', 'el-icon-Discount', 'tab', '', '/src/views/backend/user/scoreLog/index.vue', 'allow', 'none', '', 90, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(42, 41, 'button', '查看', 'user/scoreLog/index', '', '', NULL, '', '', 'deny', 'none', '', 90, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(43, 41, 'button', '添加', 'user/scoreLog/add', '', '', NULL, '', '', 'deny', 'none', '', 90, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(44, 0, 'menu_dir', '常规管理', 'routine', 'routine', 'fa fa-cogs', NULL, '', '', 'deny', 'none', '', 89, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(45, 44, 'menu', '系统配置', 'routine/config', 'routine/config', 'el-icon-Tools', 'tab', '', '/src/views/backend/routine/config/index.vue', 'allow', 'none', '', 88, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(46, 45, 'button', '查看', 'routine/config/index', '', '', NULL, '', '', 'deny', 'none', '', 88, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(47, 45, 'button', '编辑', 'routine/config/edit', '', '', NULL, '', '', 'deny', 'none', '', 88, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(48, 44, 'menu', '附件管理', 'routine/attachment', 'routine/attachment', 'fa fa-folder', 'tab', '', '/src/views/backend/routine/attachment/index.vue', 'allow', 'none', 'remark_text', 87, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(49, 48, 'button', '查看', 'routine/attachment/index', '', '', NULL, '', '', 'deny', 'none', '', 87, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(50, 48, 'button', '编辑', 'routine/attachment/edit', '', '', NULL, '', '', 'deny', 'none', '', 87, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(51, 48, 'button', '删除', 'routine/attachment/del', '', '', NULL, '', '', 'deny', 'none', '', 87, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(52, 44, 'menu', '个人资料', 'routine/adminInfo', 'routine/adminInfo', 'fa fa-user', 'tab', '', '/src/views/backend/routine/adminInfo.vue', 'allow', 'none', '', 86, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(53, 52, 'button', '查看', 'routine/adminInfo/index', '', '', NULL, '', '', 'deny', 'none', '', 86, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(54, 52, 'button', '编辑', 'routine/adminInfo/edit', '', '', NULL, '', '', 'deny', 'none', '', 86, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(55, 0, 'menu_dir', '数据安全管理', 'security', 'security', 'fa fa-shield', NULL, '', '', 'deny', 'none', '', 85, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(56, 55, 'menu', '数据回收站', 'security/dataRecycleLog', 'security/dataRecycleLog', 'fa fa-database', 'tab', '', '/src/views/backend/security/dataRecycleLog/index.vue', 'allow', 'none', '', 84, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(57, 56, 'button', '查看', 'security/dataRecycleLog/index', '', '', NULL, '', '', 'deny', 'none', '', 84, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(58, 56, 'button', '删除', 'security/dataRecycleLog/del', '', '', NULL, '', '', 'deny', 'none', '', 84, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(59, 56, 'button', '还原', 'security/dataRecycleLog/restore', '', '', NULL, '', '', 'deny', 'none', '', 84, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(60, 56, 'button', '查看详情', 'security/dataRecycleLog/info', '', '', NULL, '', '', 'deny', 'none', '', 84, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(61, 55, 'menu', '敏感数据修改记录', 'security/sensitiveDataLog', 'security/sensitiveDataLog', 'fa fa-expeditedssl', 'tab', '', '/src/views/backend/security/sensitiveDataLog/index.vue', 'allow', 'none', '', 83, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(62, 61, 'button', '查看', 'security/sensitiveDataLog/index', '', '', NULL, '', '', 'deny', 'none', '', 83, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(63, 61, 'button', '删除', 'security/sensitiveDataLog/del', '', '', NULL, '', '', 'deny', 'none', '', 83, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(64, 61, 'button', '回滚', 'security/sensitiveDataLog/rollback', '', '', NULL, '', '', 'deny', 'none', '', 83, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(65, 61, 'button', '查看详情', 'security/sensitiveDataLog/info', '', '', NULL, '', '', 'deny', 'none', '', 83, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(66, 55, 'menu', '数据回收规则管理', 'security/dataRecycle', 'security/dataRecycle', 'fa fa-database', 'tab', '', '/src/views/backend/security/dataRecycle/index.vue', 'allow', 'none', '在此定义需要回收的数据，实现数据自动统一回收', 82, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(67, 66, 'button', '查看', 'security/dataRecycle/index', '', '', NULL, '', '', 'deny', 'none', '', 82, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(68, 66, 'button', '添加', 'security/dataRecycle/add', '', '', NULL, '', '', 'deny', 'none', '', 82, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(69, 66, 'button', '编辑', 'security/dataRecycle/edit', '', '', NULL, '', '', 'deny', 'none', '', 82, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(70, 66, 'button', '删除', 'security/dataRecycle/del', '', '', NULL, '', '', 'deny', 'none', '', 82, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(71, 55, 'menu', '敏感字段规则管理', 'security/sensitiveData', 'security/sensitiveData', 'fa fa-expeditedssl', 'tab', '', '/src/views/backend/security/sensitiveData/index.vue', 'allow', 'none', '在此定义需要保护的敏感字段，随后系统将自动监听该字段的修改操作，并提供了敏感字段的修改回滚功能', 81, 'allow', '2022-07-27 01:46:21', '2022-07-27 10:31:59', NULL),
(72, 71, 'button', '查看', 'security/sensitiveData/index', '', '', NULL, '', '', 'deny', 'none', '', 81, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(73, 71, 'button', '添加', 'security/sensitiveData/add', '', '', NULL, '', '', 'deny', 'none', '', 81, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(74, 71, 'button', '编辑', 'security/sensitiveData/edit', '', '', NULL, '', '', 'deny', 'none', '', 81, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(75, 71, 'button', '删除', 'security/sensitiveData/del', '', '', NULL, '', '', 'deny', 'none', '', 81, 'allow', '2022-07-27 01:46:21', NULL, NULL),
(76, 0, 'menu', 'BuildAdmin', 'buildadmin/buildadmin', 'buildadmin', 'local-logo', 'link', 'https://doc.buildadmin.com', '', 'deny', 'none', '', 0, 'deny', '2022-07-27 01:46:21', NULL, NULL),
(77, 45, 'button', '添加', 'routine/config/add', '', '', NULL, '', '', 'deny', 'none', '', 88, 'allow', '2022-07-27 01:46:21', NULL, NULL);

INSERT INTO `admin_users` (`id`, `name`, `roles`, `password`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'test', '[\"admin\"]', '$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm', '2022-06-20 04:09:42', '2022-07-03 06:08:20', NULL),
(2, 'lose', '[\"submited\",\"published\",\"deleted\"]', 'password', '2022-07-03 06:08:50', NULL, NULL);

INSERT INTO `orders` (`id`, `story_attitude_id`, `story_id`, `user_id`, `state`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 0, 0, 'unpaid', '2022-07-17 00:52:00', '2022-07-17 00:52:00', NULL),
(2, 2, 1, 0, 'unpaid', '2022-07-17 00:54:32', '2022-07-17 00:54:32', NULL),
(3, 2, 1, 0, 'unpaid', '2022-07-17 01:00:03', '2022-07-17 01:00:03', NULL),
(4, 1, 1, 0, 'unpaid', '2022-07-17 01:02:37', '2022-07-17 01:02:37', NULL),
(5, 2, 1, 0, 'unpaid', '2022-07-17 01:07:15', '2022-07-17 01:07:15', NULL),
(6, 1, 1, 0, 'unpaid', '2022-07-17 01:09:19', '2022-07-17 01:09:19', NULL),
(7, 1, 1, 0, 'unpaid', '2022-07-17 01:10:47', '2022-07-17 01:10:47', NULL),
(8, 2, 1, 0, 'unpaid', '2022-07-17 01:12:58', '2022-07-17 01:12:58', NULL),
(9, 2, 1, 0, 'unpaid', '2022-07-17 01:17:26', '2022-07-17 01:17:26', NULL),
(10, 1, 1, 3, 'unpaid', '2022-07-17 01:19:19', '2022-07-22 08:29:31', NULL),
(11, 2, 1, 0, 'unpaid', '2022-07-17 01:20:21', '2022-07-17 01:20:21', NULL),
(12, 2, 1, 3, 'unpaid', '2022-07-17 01:22:48', '2022-07-17 01:22:48', NULL);

INSERT INTO `stories` (`id`, `title`, `code`, `content`, `state`, `rate`, `price`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '3岁女儿一眼认出爸爸石像含泪拥抱', 'L000001', '欢迎来到开拓领\n\n　　罗宾一人一马，走在漫无边际的荒原上。\n\n　　这个青年看上去二十出头，五官生得坚毅俊朗，黑色短发，黑色眼睛，眉毛笔直，鼻梁挺拔。\n\n　　他身材高大，肩膀很宽，穿着一身棕色的皮甲，腰上挂着一把长剑，看上去颇有几分英武的精气神。\n\n　　天空是蔚蓝色，万里无云，澄澈的如同一面镜子。眼前是一眼望不到头的荒草地，茁壮又茂盛的杂草有半人高，是荒原上最常见的植被。一阵微风吹过，前方的草地立刻掀起层层荡漾的波浪。\n\n　　这里是开拓领，人们口中的“北方荒原”，名义上是红曼帝国的领土，但实际上，这里是一片荒芜野蛮的未开化之地。\n\n　　罗宾现在要前往他的领地，他从未去过的领地。\n\n　　只需要向国库缴纳足够的金币，无论是贵族还是富商，甚至是平民，所有人都可以得到一张盖着帝国印章的“开拓证”，然后，你就可以前往荒原获得一片属于自己的领地了。1\n\n　　当然，帝国政府不会给你提供任何帮助，粮食、人口、军队都要你自己想办法，你得到的只是一片远离帝国边境、且鸟不拉屎的荒地而已。\n\n　　并且，生活在荒原上的蛮族土著可不认帝国的“开拓证”，如何守住这片领地，你需要自己想办法。\n\n　　开拓运动无疑是一次豪赌，倾家荡产都算是不错的结果，最惨的是把命都扔在了这片荒芜的土地上。\n\n　　但尽管如此，还是有大把大把的人费尽心机的拿到“开拓证”，然后前赴后继、义无反顾的涌向北方荒原。原因无他，拥有一片自己的土地实在太诱人了。\n\n　　不管怎么说，反正从“开拓运动”开始，整个北方荒原就渐渐变得热闹了起来。\n\n　　在前不久，罗宾也光荣的成为了开拓运动中的一员。不过他运气比较好，这张开拓证不是他花钱买的，而是天上掉下来的。\n\n　　罗宾原本是平民，他有个远的不能再远的远房叔叔做生意攒了点资产，花了大价钱搞到了一张开拓证，结果这叔叔福薄命浅，刚开拓了半年就一命呜呼的病死了。\n\n　　叔叔无儿无女，临死前他不知怎么的突然想到了罗宾这个远房侄子，或许是因为他不想自己的开拓大业半途而废，或许是因为他可怜自己这个无父无母的倒霉侄子，总之他立下了遗嘱，把领地给了罗宾。4开一局远房亲戚送业绩。\n\n　　于是就这样，罗宾莫名其妙的成为了一名男爵。\n\n　　……\n\n　　罗宾一手缰绳，一手地图，在茫茫的荒野上走了大半天的时间，终于在天黑之前赶到了他的领地。\n\n　　在一处丘陵高地上，坐落着一栋两层的黑色木楼，在帝国内部这样穷酸的房子随处可见，可是在野蛮的荒原上，这样一栋木头小楼却几乎是方圆百里内最高级别的文明建筑了。\n\n　　围绕着木楼，丘陵的坡面上零零散散的坐落着一堆房子，大概有五十户左右。这些房子简陋的可怜，就是用石头垒了四面墙，然后顶上铺上一层干草，勉强可以遮风挡雨。\n\n　　与这些石头房子相比，丘陵顶上的木楼简直是城堡级别的豪华。\n\n　　在丘陵的底部，一条蜿蜒的河流好像银丝带一样流过，河面大概有二十步宽，河水清澈见底，站在河岸边可以看见鱼儿在河底欢快的游荡。\n\n　　河岸两边水草丰茂，土地肥沃，已经被开垦出一片片农田，一些村民正弯腰在田地间劳作。\n\n　　罗宾一下子就喜欢上了这个地方，他甚至有点不敢相信，在处处贫瘠的荒原，叔叔竟然找到了这样一块风水宝地。\n\n　　他骑着马走到村子外围，看见在一条行人踏出的土路旁立着一个木头牌子，上面用帝国通用语写着：怒涛城。\n\n　　怒涛？罗宾扭头眺望了一眼远处安静祥和的河面，又看了看丘陵上这个穷苦破败的小村子。\n\n　　就这？\n\n　　怒涛城，既没有怒涛，也没有城。\n\n　　看来叔叔还挺有黑色幽默，不过不管怎样，以后这个怒涛城就是罗宾的领地了。\n\n　　这个小村子显然不习惯有外人出现，罗宾骑着马沿着土路缓缓走上丘陵，立刻就引起了村民们的注意。\n\n　　在田地间劳作的村民纷纷抬起头，赶路的村民也纷纷驻足，向罗宾投去或好奇、或警惕的目光。\n\n　　罗宾很快就发现这个村子虽然外表贫穷，但内在却很不简单，他一路上走过来，发现他的领民们竟全都是高地人！\n\n　　高地人有一个非常显著的特征，那就是长得异常高大，传说他们身上有巨人血统，男性高地人普遍身高在两米朝上，罗宾在帝国人中已经算是身材高挺，可在这里却只能跟女人比比身高。\n\n　　这些小巨人算是战斗民族，他们天生就崇尚强者和暴力，得益于他们身体上的先天优势，一个高地人单挑三四个帝国人不在话下。\n\n　　可是战争不是打群架，帝国人的文明程度对高地人是降维打击。一百年前帝国开始了对西部雪域高原的征服行动，几百个高地人部落被一扫而空。\n\n　　自此以后，高地人就成为了帝国多民族大家庭中的一员，一些部落继续留在雪域高原，而一些部落则迁移进了帝国内部。\n\n　　罗宾略微一思索，就大概猜到了他叔叔是如何把一群高地人忽悠到了荒原上。\n\n　　近些年帝国逐渐走向衰败，财政紧张、土地紧张、王权和地方紧张、大贵族之间紧张，反正啥啥都紧张。于是许多破产农民都愿意跟着小贵族来荒原上碰碰运气，反正没钱没权没地，他们烂命一条也没什么可失去了。\n\n　　对于帝国的贵族来说，高地人绝算不上理想中的优秀领民，他们除了吃饭睡觉，最擅长的事情就是打架斗殴和惹是生非，甚至很多贵族会把高地人驱逐出自己的领地，把他们当作瘟神一样赶去其他贵族的地盘。\n\n　　就连军队也很嫌弃高地人，虽然他们人高马大、力大无穷，但却毫无纪律观念，打仗只知道一股脑乱冲，有时候甚至会把己方的阵型冲垮，上演一出友军之围。\n\n　　总而言之，对于帝国人来说，高地人就是一群蒙昧的野蛮人。虽然大家都长着两只眼睛一个鼻子，但帝国人可不会把高地人当作自己的同类。\n\n　　很显然罗宾的叔叔作为一名新晋的开拓贵族，他没有太多选择领民的机会，能拉到这样一支高地部落已经算是不错的成果。\n\n　　虽然高地人有以上种种缺点，但是罗宾却欣喜若狂，因为他知道这群高地人的血脉可不简单，他们都是远古铁巨人的后裔。', 'one', 0, 0.00, '2022-07-16 01:20:28', '2022-07-16 11:35:13', NULL);

INSERT INTO `story_attitudes` (`id`, `pid`, `story_id`, `title`, `state`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 0, 1, 'ok', 'allow', '2022-07-16 10:04:02', NULL, NULL),
(2, 0, 1, 'fuck', 'allow', '2022-07-16 10:04:02', NULL, NULL);

INSERT INTO `users` (`id`, `wechat_openid`, `name`, `avatar`, `nick_name`, `password`, `account_amount`, `created_at`, `updated_at`) VALUES
(1, NULL, 'test', NULL, 'Test', '$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm', 0.00, '2022-07-13 15:01:24', '2022-07-13 15:01:35'),
(3, 'oE2tk1rxzf71mx3H2NX4rNbAAqZ4', '', '', '', '', 0.00, '2022-07-16 23:52:30', '2022-07-16 23:52:30');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;