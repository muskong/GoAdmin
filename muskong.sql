-- Database export via SQLPro (https://www.sqlprostudio.com/allapps.html)
-- Exported by a at 22-08-2022 17:54.
-- WARNING: This file may contain descructive statements such as DROPs.
-- Please ensure that you are running the script at the proper location.


-- BEGIN TABLE admin_logs
DROP TABLE IF EXISTS admin_logs;
CREATE TABLE `admin_logs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'ip地址',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求链接',
  `method` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求类型',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '资源类型',
  `param` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求参数',
  `useragent` text COLLATE utf8mb4_unicode_ci COMMENT 'User-Agent',
  `title` longtext COLLATE utf8mb4_unicode_ci COMMENT '日志',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `admin_id` (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员日志';

-- Inserting 7 rows into admin_logs
-- Insert batch #1
INSERT INTO admin_logs (id, admin_id, ip, url, method, type, param, useragent, title, created_at) VALUES
(1, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0', '登录', '2022-08-19 09:08:09'),
(2, 1, '192.168.50.35', '/admin/role/delete/4861258191949588', 'DELETE', '', '"4861258191949588"', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0', '删除角色', '2022-08-19 09:11:02'),
(3, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0', '登录', '2022-08-19 11:08:07'),
(4, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '登录', '2022-08-22 15:12:07'),
(5, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":78,"Nanoid":"7392685922348476","Type":"menu_dir","Link":"product","Path":"product/products","Icon":"\\u003cCreditCardTwoTone /\\u003e","MenuType":"sss","Remark":"产品管理","Active":"allow","Sequence":"1"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 15:19:01'),
(6, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":79,"Nanoid":"5891449357514311","ParentNanoId":"7392685922348476","Type":"menu","Title":"产品管理","Link":"product/products","Path":"product/products","Icon":"\\u003cCreditCardTwoTone /\\u003e","MenuType":"tab","Remark":"产品管理","Active":"allow","Sequence":"99"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 15:23:09'),
(7, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":80,"Nanoid":"9611653648573624","ParentNanoId":"7392685922348476","Type":"menu","Title":"卡类管理","Path":"product/card","Icon":"\\u003cCreditCardTwoTone /\\u003e","Remark":"各种卡","Active":"allow","Sequence":"99"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 16:24:00');

-- END TABLE admin_logs

-- BEGIN TABLE admin_roles
DROP TABLE IF EXISTS admin_roles;
CREATE TABLE `admin_roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nanoid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一ID',
  `parent_nanoid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '父级ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `rules` json DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '启用allow, 禁用deny',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- Inserting 15 rows into admin_roles
-- Insert batch #1
INSERT INTO admin_roles (id, nanoid, parent_nanoid, name, rules, description, state, created_at, updated_at, deleted_at) VALUES
(1, '1', '', 'admin', '[""]', '超级管理员', 'allow', '2022-07-03 06:22:53', '2022-08-11 01:45:40', NULL),
(2, '2', '1', 'admin1', NULL, '一级管理员', 'allow', '2022-08-01 08:08:54', '2022-08-10 10:46:54', NULL),
(3, 'qxJ5ZHsZoEEYjLq6cyciI', '1', 'test', '["4", "5", "6", "7", "9", "10", "11", "12", "14", "15", "16", "17", "18", "20"]', 'testeeeee', 'allow', '2022-08-11 09:50:30', '2022-08-11 09:50:30', '2022-08-11 09:50:30'),
(4, '-QvwEmNj9uO3QnlIpCnT6', '2', 'teste', '["1", "4", "5", "6", "7", "9", "10", "11", "12", "14", "15", "16", "17", "18", "20"]', 'asdfasdf', 'allow', '2022-08-11 10:00:47', '2022-08-11 10:00:47', NULL),
(5, 'o2XDvhnkbd5ffNMxo8DtS', '1', 'ew', '["46", "47", "77", "49", "50", "51", "53", "54"]', 'few', 'allow', '2022-08-11 10:03:53', '2022-08-11 03:01:03', '2022-08-11 11:01:03'),
(6, 'pqMWCbVvdO3C-9GFAOJfM', '1', 'tew', '["1", "4", "5", "6", "7", "9", "10", "11", "12", "14", "15", "16", "17", "18", "20"]', 'ewewew', 'allow', '2022-08-11 10:06:33', '2022-08-11 10:06:33', NULL),
(8, 'bnRoAcvpkS5fWA74pKOHh', '1', 'tewwwww', '["1", "4", "5", "6", "7", "9", "10", "11", "12", "14", "15", "16", "17", "18", "20"]', 'ewewewe', 'allow', '2022-08-11 10:09:15', '2022-08-11 03:00:59', '2022-08-11 11:00:59'),
(10, 'gGclxYoGZUNxx0Io4QYrL', '1', 'tewwwwwewe', '["1", "4", "5", "6", "7", "9", "10", "11", "12", "14", "15", "16", "17", "18", "20"]', 'ewewewe', 'allow', '2022-08-11 10:09:34', '2022-08-11 03:00:41', '2022-08-11 11:00:41'),
(11, '3372526697693894', '', '18221055212', '["1"]', '发烧地方', 'allow', '2022-08-11 13:29:22', '2022-08-11 13:29:22', NULL),
(12, '7532322835668198', '', 'tewe', '["1"]', 'ew', 'allow', '2022-08-11 13:31:54', '2022-08-11 05:33:38', '2022-08-11 13:33:38'),
(14, '6225617619347572', '', 'testwwww', '["1"]', 'asdf', 'allow', '2022-08-11 13:34:03', '2022-08-11 05:34:58', '2022-08-11 13:34:58'),
(15, '8889411166723482', '3372526697693894', 'tewwewda', '["1", "4", "5", "6", "7", "9", "10", "11", "12", "14", "15", "16", "17", "18", "20"]', 'bfdsbsd', 'allow', '2022-08-11 13:35:17', '2022-08-11 13:35:17', NULL),
(16, '2927997876379139', '', 'testsad', '["1"]', 'fsdfbfdsg', 'allow', '2022-08-11 13:36:09', '2022-08-11 13:36:09', NULL),
(17, '6564342879574111', '', 'test2', NULL, '', '', '2022-08-15 09:42:01', '2022-08-19 01:09:27', '2022-08-19 09:09:27'),
(18, '4861258191949588', '', 'testew2', NULL, '', '', '2022-08-15 09:43:02', '2022-08-19 01:11:02', '2022-08-19 09:11:02');

-- END TABLE admin_roles

-- BEGIN TABLE admin_rules
DROP TABLE IF EXISTS admin_rules;
CREATE TABLE `admin_rules` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `nanoid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一ID',
  `parent_nanoid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '父级ID',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'menu' COMMENT '类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由路径',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sequence` int NOT NULL DEFAULT '0' COMMENT '权重(排序)',
  `active` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态:deny=禁用,allow=启用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sequence` (`sequence`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单和权限规则表';

-- Inserting 79 rows into admin_rules
-- Insert batch #1
INSERT INTO admin_rules (id, nanoid, parent_nanoid, type, title, link, `path`, icon, remark, sequence, active, created_at, updated_at, deleted_at) VALUES
(1, '869326', '', 'menu', '控制台', 'dashboard/dashboard', 'dashboard', 'fa fa-dashboard', 'remark_text', 999, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(2, '469424', '', 'menu_dir', '权限管理', 'permission', 'permission', 'fa fa-group', '', 100, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(3, '995496', '469424', 'menu', '角色组管理', 'permission/role', 'permission/role', 'fa fa-group', '', 99, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(4, '382231', '995496', 'button', '列表', 'permission/role/list', '', '', '', 99, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(5, '187734', '995496', 'button', '创建', 'permission/role/create', '', '', '', 99, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(6, '383971', '995496', 'button', '更新', 'permission/role/update', '', '', '', 99, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(7, '762261', '995496', 'button', '删除', 'permission/role/delete', '', '', '', 99, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(8, '252513', '469424', 'menu', '管理员管理', 'permission/admin', 'permission/admin', 'el-icon-UserFilled', '', 98, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(9, '447785', '252513', 'button', '列表', 'permission/admin/list', '', '', '', 98, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(10, '833888', '252513', 'button', '创建', 'permission/admin/create', '', '', '', 98, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(11, '437416', '252513', 'button', '更新', 'permission/admin/update', '', '', '', 98, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(12, '135388', '252513', 'button', '删除', 'permission/admin/delete', '', '', '', 98, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(13, '864946', '469424', 'menu', '菜单规则管理', 'permission/rule', 'permission/rule', 'el-icon-Grid', '', 97, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(14, '272759', '864946', 'button', '列表', 'permission/rule/list', '', '', '', 97, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(15, '758594', '864946', 'button', '创建', 'permission/rule/create', '', '', '', 97, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(16, '942222', '864946', 'button', '更新', 'permission/rule/update', '', '', '', 97, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(17, '348275', '864946', 'button', '删除', 'permission/rule/delete', '', '', '', 97, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(18, '317685', '864946', 'button', '排序', 'permission/rule/sortable', '', '', '', 97, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(19, '663264', '469424', 'menu', '管理员日志管理', 'permission/adminLog', 'permission/log', 'el-icon-List', '', 96, 'allow', '2022-07-27 01:46:21', '2022-08-19 01:39:50', NULL),
(20, '156269', '663264', 'button', '列表', 'permission/adminLog/list', '', '', '', 96, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(21, '796397', '', 'menu_dir', '会员管理', 'user', 'user', 'fa fa-drivers-license', '', 95, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(22, '215312', '796397', 'menu', '会员管理', 'user/user', 'user/user', 'fa fa-user', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(23, '696251', '215312', 'button', '列表', 'user/user/list', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(24, '947511', '215312', 'button', '创建', 'user/user/create', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(25, '723213', '215312', 'button', '更新', 'user/user/update', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(26, '342833', '215312', 'button', '删除', 'user/user/delete', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(27, '387659', '796397', 'menu', '会员分组管理', 'user/group', 'user/group', 'fa fa-group', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(28, '484281', '387659', 'button', '列表', 'user/group/list', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(29, '244519', '387659', 'button', '创建', 'user/group/create', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(30, '479185', '387659', 'button', '更新', 'user/group/update', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(31, '143669', '387659', 'button', '删除', 'user/group/delete', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(32, '639121', '796397', 'menu', '会员规则管理', 'user/rule', 'user/rule', 'fa fa-th-list', '', 92, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(33, '542498', '639121', 'button', '列表', 'user/rule/list', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(34, '293761', '639121', 'button', '创建', 'user/rule/create', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(35, '335887', '639121', 'button', '更新', 'user/rule/update', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(36, '967895', '639121', 'button', '删除', 'user/rule/delete', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(37, '985743', '639121', 'button', '快速排序', 'user/rule/sortable', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(38, '942859', '796397', 'menu', '会员余额管理', 'user/moneyLog', 'user/moneyLog', 'el-icon-Money', '', 91, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(39, '192636', '942859', 'button', '列表', 'user/moneyLog/list', '', '', '', 91, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(40, '463718', '942859', 'button', '创建', 'user/moneyLog/create', '', '', '', 91, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(41, '955577', '796397', 'menu', '会员积分管理', 'user/scoreLog', 'user/scoreLog', 'el-icon-Discount', '', 90, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(42, '891947', '955577', 'button', '列表', 'user/scoreLog/list', '', '', '', 90, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(43, '258199', '955577', 'button', '创建', 'user/scoreLog/create', '', '', '', 90, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(44, '951275', '', 'menu_dir', '常规管理', 'routine', 'routine', 'fa fa-cogs', '', 89, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(45, '145924', '951275', 'menu', '系统配置', 'routine/config', 'routine/config', 'el-icon-Tools', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(46, '615994', '145924', 'button', '列表', 'routine/config/list', '', '', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(47, '624617', '145924', 'button', '更新', 'routine/config/update', '', '', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(48, '217843', '951275', 'menu', '附件管理', 'routine/attachment', 'routine/attachment', 'fa fa-folder', 'remark_text', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(49, '424991', '217843', 'button', '列表', 'routine/attachment/list', '', '', '', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(50, '977679', '217843', 'button', '更新', 'routine/attachment/update', '', '', '', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(51, '531283', '217843', 'button', '删除', 'routine/attachment/delete', '', '', '', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(52, '818444', '951275', 'menu', '个人资料', 'routine/adminInfo', 'routine/adminInfo', 'fa fa-user', '', 86, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(53, '552862', '818444', 'button', '列表', 'routine/adminInfo/list', '', '', '', 86, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(54, '895752', '818444', 'button', '更新', 'routine/adminInfo/update', '', '', '', 86, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(55, '252638', '', 'menu_dir', '数据安全管理', 'security', 'security', 'fa fa-shield', '', 85, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(56, '826733', '252638', 'menu', '数据回收站', 'security/dataRecycleLog', 'security/dataRecycleLog', 'fa fa-database', '', 84, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(57, '458777', '826733', 'button', '列表', 'security/dataRecycleLog/list', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(58, '341377', '826733', 'button', '删除', 'security/dataRecycleLog/delete', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(59, '522676', '826733', 'button', '还原', 'security/dataRecycleLog/restore', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(60, '993484', '826733', 'button', '查看详情', 'security/dataRecycleLog/info', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(61, '724847', '252638', 'menu', '敏感数据修改记录', 'security/sensitiveDataLog', 'security/sensitiveDataLog', 'fa fa-expeditedssl', '', 83, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(62, '519398', '724847', 'button', '列表', 'security/sensitiveDataLog/list', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(63, '573616', '724847', 'button', '删除', 'security/sensitiveDataLog/delete', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(64, '864146', '724847', 'button', '回滚', 'security/sensitiveDataLog/rollback', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(65, '718486', '724847', 'button', '查看详情', 'security/sensitiveDataLog/info', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(66, '745264', '252638', 'menu', '数据回收规则管理', 'security/dataRecycle', 'security/dataRecycle', 'fa fa-database', '在此定义需要回收的数据，实现数据自动统一回收', 82, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(67, '431318', '745264', 'button', '列表', 'security/dataRecycle/list', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(68, '228327', '745264', 'button', '创建', 'security/dataRecycle/create', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(69, '519558', '745264', 'button', '更新', 'security/dataRecycle/update', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(70, '646418', '745264', 'button', '删除', 'security/dataRecycle/delete', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(71, '242624', '252638', 'menu', '敏感字段规则管理', 'security/sensitiveData', 'security/sensitiveData', 'fa fa-expeditedssl', '在此定义需要保护的敏感字段，随后系统将自动监听该字段的修改操作，并提供了敏感字段的修改回滚功能', 81, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(72, '953257', '242624', 'button', '列表', 'security/sensitiveData/list', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(73, '789782', '242624', 'button', '创建', 'security/sensitiveData/create', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(74, '427915', '242624', 'button', '更新', 'security/sensitiveData/update', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(75, '691168', '242624', 'button', '删除', 'security/sensitiveData/delete', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(77, '738529', '145924', 'button', '创建', 'routine/config/create', '', '', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(78, '7392685922348476', '', 'menu_dir', '产品管理', 'product', 'product/products', '<CreditCardTwoTone />', '产品管理', 99, 'allow', '2022-08-22 15:19:01', '2022-08-22 08:10:45', NULL),
(79, '5891449357514311', '7392685922348476', 'menu', '产品管理', 'product/products', 'product/products', '<CreditCardTwoTone />', '产品管理', 99, 'allow', '2022-08-22 15:23:09', '2022-08-22 15:23:09', NULL),
(80, '9611653648573624', '7392685922348476', 'menu', '卡类管理', '', 'product/card', '<CreditCardTwoTone />', '各种卡', 99, 'allow', '2022-08-22 16:24:00', '2022-08-22 16:24:00', NULL);

-- END TABLE admin_rules

-- BEGIN TABLE admin_users
DROP TABLE IF EXISTS admin_users;
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

-- Inserting 2 rows into admin_users
-- Insert batch #1
INSERT INTO admin_users (id, name, roles, password, created_at, updated_at, deleted_at) VALUES
(1, 'test', '["admin"]', '$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm', '2022-06-20 04:09:42', '2022-07-03 06:08:20', NULL),
(2, 'lose', '["submited","published","deleted"]', 'password', '2022-07-03 06:08:50', '2022-08-15 01:21:47', '2022-08-15 09:21:47');

-- END TABLE admin_users

-- BEGIN TABLE bank_prcptcds
DROP TABLE IF EXISTS bank_prcptcds;
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

-- Table bank_prcptcds contains no data. No inserts have been genrated.
-- Inserting 0 rows into bank_prcptcds


-- END TABLE bank_prcptcds

-- BEGIN TABLE banks
DROP TABLE IF EXISTS banks;
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

-- Table banks contains no data. No inserts have been genrated.
-- Inserting 0 rows into banks


-- END TABLE banks

-- BEGIN TABLE card_orders
DROP TABLE IF EXISTS card_orders;
CREATE TABLE `card_orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一ID',
  `queue` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'queue' COMMENT '队列状态(queue等待执行, hang执行中, end执行结束)',
  `state` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单状态(hang处理中, success处理成功, error发送错误)',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `queue` (`queue`),
  KEY `state` (`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='卡订单表';

-- Table card_orders contains no data. No inserts have been genrated.
-- Inserting 0 rows into card_orders


-- END TABLE card_orders

-- BEGIN TABLE configs
DROP TABLE IF EXISTS configs;
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

-- Table configs contains no data. No inserts have been genrated.
-- Inserting 0 rows into configs


-- END TABLE configs

-- BEGIN TABLE notification_setting
DROP TABLE IF EXISTS notification_setting;
CREATE TABLE `notification_setting` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '节点名称',
  `title` varchar(255) NOT NULL COMMENT '显示名称',
  `description` varchar(255) DEFAULT NULL COMMENT '节点描述',
  `email` text COMMENT '节点邮箱发送模板',
  `email_to_user` varchar(255) DEFAULT NULL COMMENT '邮箱执行对象',
  `email_cc_user` varchar(255) DEFAULT NULL COMMENT '邮箱抄送对象',
  `email_status` varchar(255) DEFAULT 'allow' COMMENT '邮箱是否启用(allow启用, deny禁用)',
  `notification` text COMMENT '节点站内信发送模板',
  `notification_to_user` varchar(255) DEFAULT NULL COMMENT '站内信执行对象',
  `notification_cc_user` varchar(255) DEFAULT NULL COMMENT '站内信抄送对象',
  `notification_status` varchar(255) DEFAULT 'allow' COMMENT '站内信是否启用(allow启用, deny禁用)',
  `sms` text COMMENT '节点短信发送模板',
  `sms_to_user` varchar(255) DEFAULT NULL COMMENT '短信执行对象',
  `sms_cc_user` varchar(255) DEFAULT NULL COMMENT '短信抄送对象',
  `sms_status` varchar(255) DEFAULT 'allow' COMMENT '短信是否启用(allow启用, deny禁用)',
  `wechat` text COMMENT '节点企业微信发送模板',
  `wechat_to_user` varchar(255) DEFAULT NULL COMMENT '企业微信执行对象',
  `wechat_cc_user` varchar(255) DEFAULT NULL COMMENT '企业微信抄送对象',
  `wechat_status` varchar(255) DEFAULT 'allow' COMMENT '企业微信是否启用(allow启用, deny禁用)',
  `active` varchar(255) DEFAULT 'allow' COMMENT '节点状态(allow启用, deny禁用)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` int unsigned NOT NULL DEFAULT '0',
  `updated_at` datetime DEFAULT NULL,
  `updated_by` int unsigned NOT NULL DEFAULT '0',
  `deleted_at` datetime DEFAULT NULL,
  `deleted_by` int unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='通知模板设置';

-- Table notification_setting contains no data. No inserts have been genrated.
-- Inserting 0 rows into notification_setting


-- END TABLE notification_setting

-- BEGIN TABLE orders
DROP TABLE IF EXISTS orders;
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

-- Inserting 12 rows into orders
-- Insert batch #1
INSERT INTO orders (id, story_attitude_id, story_id, user_id, state, created_at, updated_at, deleted_at) VALUES
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

-- END TABLE orders

-- BEGIN TABLE product_amounts
DROP TABLE IF EXISTS product_amounts;
CREATE TABLE `product_amounts` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `amount_uuid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '产品金额唯一ID',
  `amount` decimal(10,4) NOT NULL,
  `rate` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `rate_sys` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品金额';

-- Table product_amounts contains no data. No inserts have been genrated.
-- Inserting 0 rows into product_amounts


-- END TABLE product_amounts

-- BEGIN TABLE product_cards
DROP TABLE IF EXISTS product_cards;
CREATE TABLE `product_cards` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `card_uuid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '产品唯一ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `icon_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '图片1',
  `batch` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'deny' COMMENT '是否可以批量提交卡密 allow是, deny否',
  `single` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'deny' COMMENT '是否可以单张提交卡密 allow是, deny否',
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态 allow是, deny否',
  `regularity` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '验证规则',
  `note` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `example` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '例子',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `channel` (`channel`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品卡类';

-- Table product_cards contains no data. No inserts have been genrated.
-- Inserting 0 rows into product_cards


-- END TABLE product_cards

-- BEGIN TABLE product_channels
DROP TABLE IF EXISTS product_channels;
CREATE TABLE `product_channels` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `channel_uuid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '产品渠道唯一ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '通道名称',
  `channel` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '通道(web.pc, web.mobile, api)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `channel_uuid` (`channel_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品可用渠道';

-- Table product_channels contains no data. No inserts have been genrated.
-- Inserting 0 rows into product_channels


-- END TABLE product_channels

-- BEGIN TABLE product_services
DROP TABLE IF EXISTS product_services;
CREATE TABLE `product_services` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `service_uuid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '产品服务商唯一ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `class` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '类名',
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态 allow是, deny否',
  `content` json DEFAULT NULL COMMENT '接口配置信息',
  `fields` json DEFAULT NULL COMMENT '配置字段',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `qq` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `install` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'uninstall' COMMENT 'uninstall未安装 installed 已安装 error已安装文件丢失',
  `service_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'api' COMMENT 'api 卡类接口, bank 银行接口, recharge 充值接口, tel 话单接口',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `service_uuid` (`service_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品第三方处理';

-- Table product_services contains no data. No inserts have been genrated.
-- Inserting 0 rows into product_services


-- END TABLE product_services

-- BEGIN TABLE products
DROP TABLE IF EXISTS products;
CREATE TABLE `products` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `product_card_id` int unsigned NOT NULL DEFAULT '0' COMMENT '卡类ID',
  `product_amount_id` int unsigned NOT NULL DEFAULT '0' COMMENT '金额ID',
  `product_channel_id` int unsigned NOT NULL DEFAULT '0' COMMENT '渠道ID',
  `product_service_id` int unsigned NOT NULL DEFAULT '0' COMMENT '第三方处理ID',
  `weight` int unsigned NOT NULL DEFAULT '0' COMMENT '金额在销卡中服务商占比, 权重[满分100]',
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态:deny=禁用,allow=启用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `product_card_id` (`product_card_id`),
  KEY `product_amount_id` (`product_amount_id`),
  KEY `product_channel_id` (`product_channel_id`),
  KEY `product_service_id` (`product_service_id`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品';

-- Table products contains no data. No inserts have been genrated.
-- Inserting 0 rows into products


-- END TABLE products

-- BEGIN TABLE stories
DROP TABLE IF EXISTS stories;
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

-- Inserting 1 row into stories
-- Insert batch #1
INSERT INTO stories (id, title, code, content, state, rate, price, created_at, updated_at, deleted_at) VALUES
(1, '3岁女儿一眼认出爸爸石像含泪拥抱', 'L000001', '欢迎来到开拓领\n\n　　罗宾一人一马，走在漫无边际的荒原上。\n\n　　这个青年看上去二十出头，五官生得坚毅俊朗，黑色短发，黑色眼睛，眉毛笔直，鼻梁挺拔。\n\n　　他身材高大，肩膀很宽，穿着一身棕色的皮甲，腰上挂着一把长剑，看上去颇有几分英武的精气神。\n\n　　天空是蔚蓝色，万里无云，澄澈的如同一面镜子。眼前是一眼望不到头的荒草地，茁壮又茂盛的杂草有半人高，是荒原上最常见的植被。一阵微风吹过，前方的草地立刻掀起层层荡漾的波浪。\n\n　　这里是开拓领，人们口中的“北方荒原”，名义上是红曼帝国的领土，但实际上，这里是一片荒芜野蛮的未开化之地。\n\n　　罗宾现在要前往他的领地，他从未去过的领地。\n\n　　只需要向国库缴纳足够的金币，无论是贵族还是富商，甚至是平民，所有人都可以得到一张盖着帝国印章的“开拓证”，然后，你就可以前往荒原获得一片属于自己的领地了。1\n\n　　当然，帝国政府不会给你提供任何帮助，粮食、人口、军队都要你自己想办法，你得到的只是一片远离帝国边境、且鸟不拉屎的荒地而已。\n\n　　并且，生活在荒原上的蛮族土著可不认帝国的“开拓证”，如何守住这片领地，你需要自己想办法。\n\n　　开拓运动无疑是一次豪赌，倾家荡产都算是不错的结果，最惨的是把命都扔在了这片荒芜的土地上。\n\n　　但尽管如此，还是有大把大把的人费尽心机的拿到“开拓证”，然后前赴后继、义无反顾的涌向北方荒原。原因无他，拥有一片自己的土地实在太诱人了。\n\n　　不管怎么说，反正从“开拓运动”开始，整个北方荒原就渐渐变得热闹了起来。\n\n　　在前不久，罗宾也光荣的成为了开拓运动中的一员。不过他运气比较好，这张开拓证不是他花钱买的，而是天上掉下来的。\n\n　　罗宾原本是平民，他有个远的不能再远的远房叔叔做生意攒了点资产，花了大价钱搞到了一张开拓证，结果这叔叔福薄命浅，刚开拓了半年就一命呜呼的病死了。\n\n　　叔叔无儿无女，临死前他不知怎么的突然想到了罗宾这个远房侄子，或许是因为他不想自己的开拓大业半途而废，或许是因为他可怜自己这个无父无母的倒霉侄子，总之他立下了遗嘱，把领地给了罗宾。4开一局远房亲戚送业绩。\n\n　　于是就这样，罗宾莫名其妙的成为了一名男爵。\n\n　　……\n\n　　罗宾一手缰绳，一手地图，在茫茫的荒野上走了大半天的时间，终于在天黑之前赶到了他的领地。\n\n　　在一处丘陵高地上，坐落着一栋两层的黑色木楼，在帝国内部这样穷酸的房子随处可见，可是在野蛮的荒原上，这样一栋木头小楼却几乎是方圆百里内最高级别的文明建筑了。\n\n　　围绕着木楼，丘陵的坡面上零零散散的坐落着一堆房子，大概有五十户左右。这些房子简陋的可怜，就是用石头垒了四面墙，然后顶上铺上一层干草，勉强可以遮风挡雨。\n\n　　与这些石头房子相比，丘陵顶上的木楼简直是城堡级别的豪华。\n\n　　在丘陵的底部，一条蜿蜒的河流好像银丝带一样流过，河面大概有二十步宽，河水清澈见底，站在河岸边可以看见鱼儿在河底欢快的游荡。\n\n　　河岸两边水草丰茂，土地肥沃，已经被开垦出一片片农田，一些村民正弯腰在田地间劳作。\n\n　　罗宾一下子就喜欢上了这个地方，他甚至有点不敢相信，在处处贫瘠的荒原，叔叔竟然找到了这样一块风水宝地。\n\n　　他骑着马走到村子外围，看见在一条行人踏出的土路旁立着一个木头牌子，上面用帝国通用语写着：怒涛城。\n\n　　怒涛？罗宾扭头眺望了一眼远处安静祥和的河面，又看了看丘陵上这个穷苦破败的小村子。\n\n　　就这？\n\n　　怒涛城，既没有怒涛，也没有城。\n\n　　看来叔叔还挺有黑色幽默，不过不管怎样，以后这个怒涛城就是罗宾的领地了。\n\n　　这个小村子显然不习惯有外人出现，罗宾骑着马沿着土路缓缓走上丘陵，立刻就引起了村民们的注意。\n\n　　在田地间劳作的村民纷纷抬起头，赶路的村民也纷纷驻足，向罗宾投去或好奇、或警惕的目光。\n\n　　罗宾很快就发现这个村子虽然外表贫穷，但内在却很不简单，他一路上走过来，发现他的领民们竟全都是高地人！\n\n　　高地人有一个非常显著的特征，那就是长得异常高大，传说他们身上有巨人血统，男性高地人普遍身高在两米朝上，罗宾在帝国人中已经算是身材高挺，可在这里却只能跟女人比比身高。\n\n　　这些小巨人算是战斗民族，他们天生就崇尚强者和暴力，得益于他们身体上的先天优势，一个高地人单挑三四个帝国人不在话下。\n\n　　可是战争不是打群架，帝国人的文明程度对高地人是降维打击。一百年前帝国开始了对西部雪域高原的征服行动，几百个高地人部落被一扫而空。\n\n　　自此以后，高地人就成为了帝国多民族大家庭中的一员，一些部落继续留在雪域高原，而一些部落则迁移进了帝国内部。\n\n　　罗宾略微一思索，就大概猜到了他叔叔是如何把一群高地人忽悠到了荒原上。\n\n　　近些年帝国逐渐走向衰败，财政紧张、土地紧张、王权和地方紧张、大贵族之间紧张，反正啥啥都紧张。于是许多破产农民都愿意跟着小贵族来荒原上碰碰运气，反正没钱没权没地，他们烂命一条也没什么可失去了。\n\n　　对于帝国的贵族来说，高地人绝算不上理想中的优秀领民，他们除了吃饭睡觉，最擅长的事情就是打架斗殴和惹是生非，甚至很多贵族会把高地人驱逐出自己的领地，把他们当作瘟神一样赶去其他贵族的地盘。\n\n　　就连军队也很嫌弃高地人，虽然他们人高马大、力大无穷，但却毫无纪律观念，打仗只知道一股脑乱冲，有时候甚至会把己方的阵型冲垮，上演一出友军之围。\n\n　　总而言之，对于帝国人来说，高地人就是一群蒙昧的野蛮人。虽然大家都长着两只眼睛一个鼻子，但帝国人可不会把高地人当作自己的同类。\n\n　　很显然罗宾的叔叔作为一名新晋的开拓贵族，他没有太多选择领民的机会，能拉到这样一支高地部落已经算是不错的成果。\n\n　　虽然高地人有以上种种缺点，但是罗宾却欣喜若狂，因为他知道这群高地人的血脉可不简单，他们都是远古铁巨人的后裔。', 'one', 0, 0, '2022-07-16 01:20:28', '2022-07-16 11:35:13', NULL);

-- END TABLE stories

-- BEGIN TABLE story_attitudes
DROP TABLE IF EXISTS story_attitudes;
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

-- Inserting 2 rows into story_attitudes
-- Insert batch #1
INSERT INTO story_attitudes (id, pid, story_id, title, state, created_at, updated_at, deleted_at) VALUES
(1, 0, 1, 'ok', 'allow', '2022-07-16 10:04:02', NULL, NULL),
(2, 0, 1, 'fuck', 'allow', '2022-07-16 10:04:02', NULL, NULL);

-- END TABLE story_attitudes

-- BEGIN TABLE story_progress
DROP TABLE IF EXISTS story_progress;
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

-- Table story_progress contains no data. No inserts have been genrated.
-- Inserting 0 rows into story_progress


-- END TABLE story_progress

-- BEGIN TABLE tmp
DROP TABLE IF EXISTS tmp;
CREATE TABLE `tmp` (
  `id` int NOT NULL AUTO_INCREMENT,
  `string` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nanoid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Inserting 99 rows into tmp
-- Insert batch #1
INSERT INTO tmp (id, string, nanoid) VALUES
(1, 'rKxtBf', '869326'),
(2, 'WiCMIs', '469424'),
(3, 'HuYsuH', '995496'),
(4, 'tXFadm', '382231'),
(5, 'BUCGLZ', '187734'),
(6, 'TtdGKj', '383971'),
(7, 'mNlSmc', '762261'),
(8, 'TTFDFR', '252513'),
(9, 'LEQpum', '447785'),
(10, 'AsDhyM', '833888'),
(11, 'ByMbhw', '437416'),
(12, 'lQzDXM', '135388'),
(13, 'FRzWft', '864946'),
(14, 'QpsSMW', '272759'),
(15, 'CfGCUF', '758594'),
(16, 'VdEVuI', '942222'),
(17, 'GWyPdA', '348275'),
(18, 'picPvg', '317685'),
(19, 'HaBlcW', '663264'),
(20, 'RUHmiq', '156269'),
(21, 'zqjASF', '796397'),
(22, 'aauRhz', '215312'),
(23, 'wRNgvQ', '696251'),
(24, 'HryVXs', '947511'),
(25, 'eyiNeN', '723213'),
(26, 'IRwQHU', '342833'),
(27, 'qBhcwf', '387659'),
(28, 'EnsnhD', '484281'),
(29, 'ACJxVz', '244519'),
(30, 'eFWHWH', '479185'),
(31, 'MIhFCp', '143669'),
(32, 'gTtDib', '639121'),
(33, 'FBAdTF', '542498'),
(34, 'TFFMSR', '293761'),
(35, 'knuSCM', '335887'),
(36, 'nUnMaI', '967895'),
(37, 'GLPzYh', '985743'),
(38, 'HxlGxM', '942859'),
(39, 'cVwfVZ', '192636'),
(40, 'XZgnaQ', '463718'),
(41, 'cTFdqY', '955577'),
(42, 'bjSjXT', '891947'),
(43, 'TagdDC', '258199'),
(44, 'QJgUrI', '951275'),
(45, 'WLNkFB', '145924'),
(46, 'lmkcXA', '615994'),
(47, 'axArxb', '624617'),
(48, 'gGyVqt', '217843'),
(49, 'eifhwD', '424991'),
(50, 'MPKgtd', '977679'),
(51, 'jSZRhD', '531283'),
(52, 'EscWMz', '818444'),
(53, 'gnMjuz', '552862'),
(54, 'qzWvmK', '895752'),
(55, 'clyZAb', '252638'),
(56, 'sUJsXv', '826733'),
(57, 'TRYCSa', '458777'),
(58, 'ekGDkt', '341377'),
(59, 'TUjJLT', '522676'),
(60, 'LCEIxR', '993484'),
(61, 'czPuea', '724847'),
(62, 'TLUTSD', '519398'),
(63, 'SZecaz', '573616'),
(64, 'rKQFee', '864146'),
(65, 'zlBwYp', '718486'),
(66, 'dkmwzl', '745264'),
(67, 'JZMbvb', '431318'),
(68, 'bIFWnG', '228327'),
(69, 'yQYgnX', '519558'),
(70, 'UzhIcP', '646418'),
(71, 'TKAafy', '242624'),
(72, 'IyRYEa', '953257'),
(73, 'dWxBnF', '789782'),
(74, 'CpLkBA', '427915'),
(75, 'TaeLpG', '691168'),
(76, 'ZyNFdu', '341852'),
(77, 'vVfGFz', '738529'),
(78, 'BIFTie', '872579'),
(79, 'SuJMyD', '277353'),
(80, 'LkdzZC', '368464'),
(81, 'sjnZcd', '119673'),
(82, 'NNDyZd', '848732'),
(83, 'cepLkf', '477267'),
(84, 'pxLECL', '556662'),
(85, 'VRAcXj', '127242'),
(86, 'kvnGYM', '875966'),
(87, 'UpSwjb', '426892'),
(88, 'nCbgwl', '269579'),
(89, 'FuxjMQ', '694283'),
(90, 'qylgsb', '216333'),
(91, 'bgGdbh', '936263'),
(92, 'MvAlHk', '885981'),
(93, 'ZcXnnj', '266232'),
(94, 'YdSLnI', '435581'),
(95, 'mVikwz', '665161'),
(96, 'rFwyVT', '729488'),
(97, 'qSjbRY', '439341'),
(98, 'GvgkJe', '338611'),
(99, 'jpQzzw', '838516');

-- END TABLE tmp

-- BEGIN TABLE user_accounts
DROP TABLE IF EXISTS user_accounts;
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

-- Table user_accounts contains no data. No inserts have been genrated.
-- Inserting 0 rows into user_accounts


-- END TABLE user_accounts

-- BEGIN TABLE user_auths
DROP TABLE IF EXISTS user_auths;
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

-- Table user_auths contains no data. No inserts have been genrated.
-- Inserting 0 rows into user_auths


-- END TABLE user_auths

-- BEGIN TABLE user_banks
DROP TABLE IF EXISTS user_banks;
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

-- Table user_banks contains no data. No inserts have been genrated.
-- Inserting 0 rows into user_banks


-- END TABLE user_banks

-- BEGIN TABLE users
DROP TABLE IF EXISTS users;
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

-- Inserting 2 rows into users
-- Insert batch #1
INSERT INTO users (id, wechat_openid, name, avatar, nick_name, password, account_amount, created_at, updated_at) VALUES
(1, NULL, 'test', NULL, 'Test', '$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm', 0, '2022-07-13 15:01:24', '2022-07-13 15:01:35'),
(3, 'oE2tk1rxzf71mx3H2NX4rNbAAqZ4', '', '', '', '', 0, '2022-07-16 23:52:30', '2022-07-16 23:52:30');

-- END TABLE users
