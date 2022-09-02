-- Database export via SQLPro (https://www.sqlprostudio.com/allapps.html)
-- Exported by a at 01-09-2022 19:19.
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
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员日志';

-- Inserting 32 rows into admin_logs
-- Insert batch #1
INSERT INTO admin_logs (id, admin_id, ip, url, method, type, param, useragent, title, created_at) VALUES
(1, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0', '登录', '2022-08-19 09:08:09'),
(2, 1, '192.168.50.35', '/admin/role/delete/4861258191949588', 'DELETE', '', '"4861258191949588"', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0', '删除角色', '2022-08-19 09:11:02'),
(3, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0', '登录', '2022-08-19 11:08:07'),
(4, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '登录', '2022-08-22 15:12:07'),
(5, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":78,"Nanoid":"7392685922348476","Type":"menu_dir","Link":"product","Path":"product/products","Icon":"\\u003cCreditCardTwoTone /\\u003e","MenuType":"sss","Remark":"产品管理","Active":"allow","Sequence":"1"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 15:19:01'),
(6, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":79,"Nanoid":"5891449357514311","ParentNanoId":"7392685922348476","Type":"menu","Title":"产品管理","Link":"product/products","Path":"product/products","Icon":"\\u003cCreditCardTwoTone /\\u003e","MenuType":"tab","Remark":"产品管理","Active":"allow","Sequence":"99"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 15:23:09'),
(7, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":80,"Nanoid":"9611653648573624","ParentNanoId":"7392685922348476","Type":"menu","Title":"卡类管理","Path":"product/card","Icon":"\\u003cCreditCardTwoTone /\\u003e","Remark":"各种卡","Active":"allow","Sequence":"99"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 16:24:00'),
(8, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":81,"Nanoid":"7666542485486136","ParentNanoId":"7392685922348476","Type":"menu","Title":"金额管理","Path":"product/amount","Icon":"\\u003cMoneyCollectTwoTone /\\u003e","Remark":"金额管理","Active":"allow","Sequence":"99"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 18:22:11'),
(9, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":82,"Nanoid":"8221982471489842","ParentNanoId":"7392685922348476","Type":"menu","Title":"渠道管理","Path":"product/channel","Icon":"\\u003cGoldTwoTone /\\u003e","Remark":"渠道","Active":"allow","Sequence":"99"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 18:23:30'),
(10, 1, '192.168.50.35', '/admin/rule/create', 'POST', 'application/json', '{"Id":83,"Nanoid":"9678964433246326","ParentNanoId":"7392685922348476","Type":"menu","Title":"服务商管理","Path":"product/service","Icon":"\\u003cCustomerServiceTwoTone /\\u003e","Remark":"服务商","Active":"allow","Sequence":"99"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增权限节点', '2022-08-22 18:24:21'),
(11, 1, '192.168.50.35', '/admin/productAmount/create', 'POST', 'application/json', '{"ID":0,"Id":1,"Amount":"100","Rate":"0.94","RateSys":"0.98"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品金额', '2022-08-22 18:27:12'),
(12, 1, '192.168.50.35', '/admin/productAmount/create', 'POST', 'application/json', '{"ID":0,"Id":2,"Amount":"1000","Rate":"0.94","RateSys":"0.98"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品金额', '2022-08-22 18:27:25'),
(13, 1, '192.168.50.35', '/admin/productAmount/create', 'POST', 'application/json', '{"ID":0,"Id":3,"Amount":"200","Rate":"0.94","RateSys":"0.98"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品金额', '2022-08-22 18:27:36'),
(14, 1, '192.168.50.35', '/admin/productAmount/create', 'POST', 'application/json', '{"ID":0,"Id":4,"Amount":"300","Rate":"0.94","RateSys":"0.98"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品金额', '2022-08-22 18:28:36'),
(15, 1, '192.168.50.35', '/admin/productChannel/create', 'POST', 'application/json', '{"ID":0,"Id":1,"Title":"手机","Channel":"web.mobile"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品渠道', '2022-08-22 18:31:12'),
(16, 1, '192.168.50.35', '/admin/productChannel/create', 'POST', 'application/json', '{"ID":0,"Id":2,"Title":"电脑","Channel":"web.pc"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品渠道', '2022-08-22 18:31:26'),
(17, 1, '192.168.50.35', '/admin/productChannel/create', 'POST', 'application/json', '{"ID":0,"Id":3,"Title":"接口","Channel":"api"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品渠道', '2022-08-22 18:31:40'),
(18, 1, '192.168.50.35', '/admin/productCard/create', 'POST', 'application/json', '{"ID":0,"Id":1,"Title":"京东E卡","IconUrl":"url","Batch":"allow","Single":"allow","Status":"allow","Regularity":"rage","Note":"京东E卡","Example":"1234-1234-1234-1234"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54', '新增产品卡类', '2022-08-22 18:42:27'),
(19, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '登录', '2022-08-23 18:40:39'),
(20, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '登录', '2022-08-24 12:31:57'),
(21, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '登录', '2022-08-25 17:25:56'),
(22, 1, '192.168.50.35', '/admin/productAmount/create', 'POST', 'application/json', '{"ID":0,"Id":5}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '新增产品金额', '2022-08-26 13:45:53'),
(23, 1, '192.168.50.35', '/admin/productAmount/create', 'POST', 'application/json', '{"ID":0,"Id":6}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '新增产品金额', '2022-08-26 14:49:19'),
(24, 1, '192.168.50.35', '/admin/productService/install', 'POST', 'application/json', '{"PayAccount":"4","PayKey":"2","PayName":"3","PayPid":"1","SearchUrl":"6","SendUrl":"5"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '新增产品服务商', '2022-08-26 15:37:52'),
(25, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '登录', '2022-08-29 08:56:32'),
(26, 1, '192.168.50.35', '/admin/productAmount/delete/6', 'DELETE', '', '6', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '删除产品金额', '2022-08-29 09:12:58'),
(27, 1, '192.168.50.35', '/admin/productAmount/delete/5', 'DELETE', '', '5', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '删除产品金额', '2022-08-29 09:13:19'),
(28, 1, '192.168.50.35', '/admin/product/create', 'POST', 'application/json', '{"ID":0,"UpdatedAt":"0001-01-01T00:00:00Z","Id":1,"ProductCardId":1,"ProductAmountId":4,"ProductChannelId":3,"ProductServiceId":1,"Weight":11,"Status":"allow","Card":{"ID":0},"Amount":{"ID":0},"Channel":{"ID":0},"Service":{"ID":0}}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '新增产品', '2022-08-29 09:25:59'),
(29, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.63', '登录', '2022-08-30 09:27:36'),
(30, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.70', '登录', '2022-08-31 18:46:47'),
(31, 1, '192.168.50.35', '/admin/userGroup/create', 'POST', 'application/json', '{"ID":1,"GroupUuid":"twaAwNRZAxnHCVHrIhCuTbyeljquHJ","Title":"test","Content":"342"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.70', '新增用户组', '2022-09-01 17:33:36'),
(32, 0, '192.168.50.35', '/admin/admin/login', 'POST', 'application/json', '{"Id":1,"Name":"test","Password":"$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm","Roles":["admin"],"CreatedAt":"2022-06-20 04:09:42","UpdatedAt":"2022-07-03 06:08:20"}', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.102 Safari/537.36 Edg/104.0.1293.70', '登录', '2022-09-01 18:58:05');

-- END TABLE admin_logs

-- BEGIN TABLE admin_roles
DROP TABLE IF EXISTS admin_roles;
CREATE TABLE `admin_roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nanoid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一ID',
  `parent_nanoid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '父级ID',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `rules` json DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '启用allow, 禁用deny',
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
  `nanoid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一ID',
  `parent_nanoid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '父级ID',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'menu' COMMENT '类型:menu_dir=菜单目录,menu=菜单项,button=页面按钮',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `link` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由路径',
  `icon` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sequence` int NOT NULL DEFAULT '0' COMMENT '权重(排序)',
  `active` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态:deny=禁用,allow=启用',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sequence` (`sequence`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单和权限规则表';

-- Inserting 82 rows into admin_rules
-- Insert batch #1
INSERT INTO admin_rules (id, nanoid, parent_nanoid, type, title, link, `path`, icon, remark, sequence, active, created_at, updated_at, deleted_at) VALUES
(1, '869326', '', 'menu', '控制台', 'dashboard/dashboard', 'dashboard', 'fa fa-dashboard', 'remark_text', 999, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(2, '469424', '', 'menu_dir', '权限管理', 'permission', 'permission', 'fa fa-group', '', 1, 'allow', '2022-07-27 01:46:21', '2022-09-01 10:22:50', NULL),
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
(21, '796397', '', 'menu_dir', '用户管理', 'user', 'user', 'fa fa-drivers-license', '', 95, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(22, '215312', '796397', 'menu', '用户管理', 'user/user', 'user/user', 'fa fa-user', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(23, '696251', '215312', 'button', '列表', 'user/user/list', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(24, '947511', '215312', 'button', '创建', 'user/user/create', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(25, '723213', '215312', 'button', '更新', 'user/user/update', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(26, '342833', '215312', 'button', '删除', 'user/user/delete', '', '', '', 94, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(27, '387659', '796397', 'menu', '用户分组管理', 'user/group', 'user/group', 'fa fa-group', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(28, '484281', '387659', 'button', '列表', 'user/group/list', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(29, '244519', '387659', 'button', '创建', 'user/group/create', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(30, '479185', '387659', 'button', '更新', 'user/group/update', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(31, '143669', '387659', 'button', '删除', 'user/group/delete', '', '', '', 93, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(32, '639121', '796397', 'menu', '用户认证管理', 'user/verified', 'user/verified', 'fa fa-th-list', '', 92, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(33, '542498', '639121', 'button', '列表', 'user/verified/list', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(34, '293761', '639121', 'button', '创建', 'user/verified/create', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(35, '335887', '639121', 'button', '更新', 'user/verified/update', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(36, '967895', '639121', 'button', '删除', 'user/verified/delete', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(37, '985743', '639121', 'button', '快速排序', 'user/verified/sortable', '', '', '', 92, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(38, '942859', '796397', 'menu', '用户交易记录', 'user/account', 'user/account', 'el-icon-Money', '', 91, 'allow', '2022-07-27 01:46:21', '2022-09-01 08:51:23', NULL),
(39, '192636', '942859', 'button', '列表', 'user/account/list', '', '', '', 91, 'allow', '2022-07-27 01:46:21', '2022-09-01 08:51:23', NULL),
(40, '463718', '942859', 'button', '创建', 'user/account/create', '', '', '', 91, 'allow', '2022-07-27 01:46:21', '2022-09-01 08:51:23', NULL),
(41, '955577', '796397', 'menu', '用户银行卡管理', 'user/bank', 'user/bank', 'el-icon-Discount', '', 90, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(42, '891947', '955577', 'button', '列表', 'user/bank/list', '', '', '', 90, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(43, '258199', '955577', 'button', '创建', 'user/bank/create', '', '', '', 90, 'allow', '2022-07-27 01:46:21', '2022-09-01 09:42:39', NULL),
(44, '951275', '', 'menu_dir', '常规管理', 'routine', 'routine', 'fa fa-cogs', '', 89, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:07', NULL),
(45, '145924', '951275', 'menu', '系统配置', 'routine/config', 'routine/config', 'el-icon-Tools', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(46, '615994', '145924', 'button', '列表', 'routine/config/list', '', '', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(47, '624617', '145924', 'button', '更新', 'routine/config/update', '', '', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(48, '217843', '951275', 'menu', '附件管理', 'routine/attachment', 'routine/attachment', 'fa fa-folder', 'remark_text', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(49, '424991', '217843', 'button', '列表', 'routine/attachment/list', '', '', '', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(50, '977679', '217843', 'button', '更新', 'routine/attachment/update', '', '', '', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(51, '531283', '217843', 'button', '删除', 'routine/attachment/delete', '', '', '', 87, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(52, '818444', '951275', 'menu', '个人资料', 'routine/adminInfo', 'routine/adminInfo', 'fa fa-user', '', 86, 'allow', '2022-07-27 01:46:21', '2022-09-01 10:30:02', '2022-09-01 10:30:02'),
(53, '552862', '818444', 'button', '列表', 'routine/adminInfo/list', '', '', '', 86, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(54, '895752', '818444', 'button', '更新', 'routine/adminInfo/update', '', '', '', 86, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(55, '252638', '', 'menu_dir', '数据安全管理', 'security', 'security', 'fa fa-shield', '', 85, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:09:18', '2022-09-01 06:09:18'),
(56, '826733', '252638', 'menu', '数据回收站', 'security/dataRecycleLog', 'security/dataRecycleLog', 'fa fa-database', '', 84, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:09:23', '2022-09-01 06:09:23'),
(57, '458777', '826733', 'button', '列表', 'security/dataRecycleLog/list', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(58, '341377', '826733', 'button', '删除', 'security/dataRecycleLog/delete', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(59, '522676', '826733', 'button', '还原', 'security/dataRecycleLog/restore', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(60, '993484', '826733', 'button', '查看详情', 'security/dataRecycleLog/info', '', '', '', 84, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(61, '724847', '252638', 'menu', '敏感数据修改记录', 'security/sensitiveDataLog', 'security/sensitiveDataLog', 'fa fa-expeditedssl', '', 83, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:09:23', '2022-09-01 06:09:23'),
(62, '519398', '724847', 'button', '列表', 'security/sensitiveDataLog/list', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(63, '573616', '724847', 'button', '删除', 'security/sensitiveDataLog/delete', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(64, '864146', '724847', 'button', '回滚', 'security/sensitiveDataLog/rollback', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(65, '718486', '724847', 'button', '查看详情', 'security/sensitiveDataLog/info', '', '', '', 83, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(66, '745264', '252638', 'menu', '数据回收规则管理', 'security/dataRecycle', 'security/dataRecycle', 'fa fa-database', '在此定义需要回收的数据，实现数据自动统一回收', 82, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:09:23', '2022-09-01 06:09:23'),
(67, '431318', '745264', 'button', '列表', 'security/dataRecycle/list', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(68, '228327', '745264', 'button', '创建', 'security/dataRecycle/create', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(69, '519558', '745264', 'button', '更新', 'security/dataRecycle/update', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(70, '646418', '745264', 'button', '删除', 'security/dataRecycle/delete', '', '', '', 82, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(71, '242624', '252638', 'menu', '敏感字段规则管理', 'security/sensitiveData', 'security/sensitiveData', 'fa fa-expeditedssl', '在此定义需要保护的敏感字段，随后系统将自动监听该字段的修改操作，并提供了敏感字段的修改回滚功能', 81, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:09:23', '2022-09-01 06:09:23'),
(72, '953257', '242624', 'button', '列表', 'security/sensitiveData/list', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(73, '789782', '242624', 'button', '创建', 'security/sensitiveData/create', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(74, '427915', '242624', 'button', '更新', 'security/sensitiveData/update', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(75, '691168', '242624', 'button', '删除', 'security/sensitiveData/delete', '', '', '', 81, 'allow', '2022-07-27 01:46:21', '2022-09-01 06:11:10', '2022-09-01 06:11:10'),
(77, '738529', '145924', 'button', '创建', 'routine/config/create', '', '', '', 88, 'allow', '2022-07-27 01:46:21', '2022-08-12 07:03:22', NULL),
(78, '7392685922348476', '', 'menu_dir', '产品管理', 'product', 'product/products', '<CreditCardTwoTone />', '产品管理', 99, 'allow', '2022-08-22 15:19:01', '2022-08-22 08:10:45', NULL),
(79, '5891449357514311', '7392685922348476', 'menu', '产品管理', 'product/products', 'product/products', '<CreditCardTwoTone />', '产品管理', 99, 'allow', '2022-08-22 15:23:09', '2022-08-22 15:23:09', NULL),
(80, '9611653648573624', '7392685922348476', 'menu', '卡类管理', '', 'product/card', '<CreditCardTwoTone />', '各种卡', 98, 'allow', '2022-08-22 16:24:00', '2022-09-01 10:28:12', NULL),
(81, '7666542485486136', '7392685922348476', 'menu', '金额管理', '', 'product/amount', '<MoneyCollectTwoTone />', '金额管理', 97, 'allow', '2022-08-22 18:22:11', '2022-09-01 10:28:12', NULL),
(82, '8221982471489842', '7392685922348476', 'menu', '渠道管理', '', 'product/channel', '<GoldTwoTone />', '渠道', 96, 'allow', '2022-08-22 18:23:30', '2022-09-01 10:28:12', NULL),
(83, '9678964433246326', '7392685922348476', 'menu', '服务商管理', '', 'product/service', '<CustomerServiceTwoTone />', '服务商', 95, 'allow', '2022-08-22 18:24:21', '2022-09-01 10:28:12', NULL);

-- END TABLE admin_rules

-- BEGIN TABLE admin_users
DROP TABLE IF EXISTS admin_users;
CREATE TABLE `admin_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `roles` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
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
  `bankname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `bank_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '银行编号',
  `prcptcd` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '行号',
  `city_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '所属地区编号',
  `province` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
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
  `abbr` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` int NOT NULL DEFAULT '1',
  `bankname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `logo` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `logo1` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `procode` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='银行';

-- Table banks contains no data. No inserts have been genrated.
-- Inserting 0 rows into banks


-- END TABLE banks

-- BEGIN TABLE configs
DROP TABLE IF EXISTS configs;
CREATE TABLE `configs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
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
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '节点描述',
  `email` text COMMENT '节点邮箱发送模板',
  `email_to_user` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱执行对象',
  `email_cc_user` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱抄送对象',
  `email_status` varchar(255) DEFAULT 'allow' COMMENT '邮箱是否启用(allow启用, deny禁用)',
  `notification` text COMMENT '节点站内信发送模板',
  `notification_to_user` varchar(255) NOT NULL DEFAULT '' COMMENT '站内信执行对象',
  `notification_cc_user` varchar(255) NOT NULL DEFAULT '' COMMENT '站内信抄送对象',
  `notification_status` varchar(255) DEFAULT 'allow' COMMENT '站内信是否启用(allow启用, deny禁用)',
  `sms` text COMMENT '节点短信发送模板',
  `sms_to_user` varchar(255) NOT NULL DEFAULT '' COMMENT '短信执行对象',
  `sms_cc_user` varchar(255) NOT NULL DEFAULT '' COMMENT '短信抄送对象',
  `sms_status` varchar(255) DEFAULT 'allow' COMMENT '短信是否启用(allow启用, deny禁用)',
  `wechat` text COMMENT '节点企业微信发送模板',
  `wechat_to_user` varchar(255) NOT NULL DEFAULT '' COMMENT '企业微信执行对象',
  `wechat_cc_user` varchar(255) NOT NULL DEFAULT '' COMMENT '企业微信抄送对象',
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
  `order_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '唯一ID',
  `product_id` int unsigned NOT NULL DEFAULT '0' COMMENT '产品ID',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '产品ID',
  `external_id` int unsigned NOT NULL DEFAULT '0' COMMENT '外部ID',
  `channel` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '渠道(web.pc, web.mobile, api)',
  `queue` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'queue' COMMENT '队列状态(queue等待执行, hang执行中, end执行结束)',
  `state` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '订单状态(hang处理中, success处理成功, error发送错误)',
  `card_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '卡号',
  `card_password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `card_cvv` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'cvv',
  `result` json DEFAULT NULL COMMENT '返回数据',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `order_number` (`order_number`),
  KEY `product_id` (`product_id`),
  KEY `queue` (`queue`),
  KEY `state` (`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='卡订单表';

-- Table orders contains no data. No inserts have been genrated.
-- Inserting 0 rows into orders


-- END TABLE orders

-- BEGIN TABLE product_amounts
DROP TABLE IF EXISTS product_amounts;
CREATE TABLE `product_amounts` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `amount_uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '产品金额唯一ID',
  `amount` decimal(10,4) NOT NULL,
  `rate` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `rate_sys` decimal(10,4) NOT NULL DEFAULT '0.0000',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品金额';

-- Inserting 6 rows into product_amounts
-- Insert batch #1
INSERT INTO product_amounts (id, amount_uuid, amount, rate, rate_sys, created_at, updated_at, deleted_at) VALUES
(1, '', 100, 0.94, 0.98, '2022-08-22 18:27:12', '2022-08-22 18:27:12', NULL),
(2, '', 1000, 0.94, 0.98, '2022-08-22 18:27:25', '2022-08-22 18:27:25', NULL),
(3, '', 200, 0.94, 0.98, '2022-08-22 18:27:36', '2022-08-22 18:27:36', NULL),
(4, '', 300, 0.94, 0.98, '2022-08-22 18:28:36', '2022-08-22 18:28:36', NULL),
(5, '', 0, 0, 0, '2022-08-26 13:45:53', '2022-08-29 01:13:19', '2022-08-29 09:13:18'),
(6, '', 0, 0, 0, '2022-08-26 14:49:19', '2022-08-29 01:12:58', '2022-08-29 09:12:58');

-- END TABLE product_amounts

-- BEGIN TABLE product_cards
DROP TABLE IF EXISTS product_cards;
CREATE TABLE `product_cards` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `card_uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '产品唯一ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `icon_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图片1',
  `batch` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'deny' COMMENT '是否可以批量提交卡密 allow是, deny否',
  `single` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'deny' COMMENT '是否可以单张提交卡密 allow是, deny否',
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态 allow是, deny否',
  `regularity` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '验证规则',
  `note` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `example` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '例子',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `card_uuid` (`card_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品卡类';

-- Inserting 1 row into product_cards
-- Insert batch #1
INSERT INTO product_cards (id, card_uuid, title, icon_url, batch, single, status, regularity, note, example, created_at, updated_at, deleted_at) VALUES
(1, '', '京东E卡', 'url', 'allow', 'allow', 'allow', 'rage', '京东E卡', '1234-1234-1234-1234', '2022-08-22 18:42:27', '2022-08-22 18:42:27', NULL);

-- END TABLE product_cards

-- BEGIN TABLE product_channels
DROP TABLE IF EXISTS product_channels;
CREATE TABLE `product_channels` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `channel_uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '产品渠道唯一ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '通道名称',
  `channel` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '通道(web.pc, web.mobile, api)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `channel_uuid` (`channel_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品可用渠道';

-- Inserting 3 rows into product_channels
-- Insert batch #1
INSERT INTO product_channels (id, channel_uuid, title, channel, created_at, updated_at, deleted_at) VALUES
(1, '', '手机', 'web.mobile', '2022-08-22 18:31:12', '2022-08-22 18:31:12', NULL),
(2, '', '电脑', 'web.pc', '2022-08-22 18:31:26', '2022-08-22 18:31:26', NULL),
(3, '', '接口', 'api', '2022-08-22 18:31:40', '2022-08-22 18:31:40', NULL);

-- END TABLE product_channels

-- BEGIN TABLE product_services
DROP TABLE IF EXISTS product_services;
CREATE TABLE `product_services` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '产品服务商唯一ID',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标题',
  `class` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '类名',
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'allow' COMMENT '状态 allow是, deny否',
  `content` json DEFAULT NULL COMMENT '接口配置信息',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'api' COMMENT 'api 卡类接口, bank 银行接口, recharge 充值接口, tel 话单接口',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品第三方处理';

-- Inserting 2 rows into product_services
-- Insert batch #1
INSERT INTO product_services (id, uuid, title, class, status, content, type, created_at, updated_at, deleted_at) VALUES
(1, 'GXVSemcZteNWbjpL', '收卡云', 'ShouKaYun', 'deny', '{"PayKey": "2", "PayPid": "1", "PayName": "3", "SendUrl": "5", "SearchUrl": "6", "PayAccount": "4"}', 'api', '2022-08-26 15:35:42', '2022-08-31 10:47:26', '2022-08-31 18:47:26'),
(2, 'petRTIcLDndFPHzX', '收卡云', 'ShouKaYun', 'deny', '{"PayKey": "2", "PayPid": "1", "PayName": "3", "SendUrl": "5", "SearchUrl": "6", "PayAccount": "4"}', 'api', '2022-08-26 15:37:52', '2022-08-26 08:53:05', '2022-08-26 16:53:05');

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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='产品';

-- Inserting 1 row into products
-- Insert batch #1
INSERT INTO products (id, product_card_id, product_amount_id, product_channel_id, product_service_id, weight, status, created_at, deleted_at) VALUES
(1, 1, 4, 3, 1, 11, 'allow', '2022-08-29 09:25:59', NULL);

-- END TABLE products

-- BEGIN TABLE tmp
DROP TABLE IF EXISTS tmp;
CREATE TABLE `tmp` (
  `id` int NOT NULL AUTO_INCREMENT,
  `string` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `nanoid` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
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
  `user_uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `before` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '变更前金额',
  `change` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '变更金额',
  `after` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '变更后金额',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '变更说明',
  `table` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '变更对应的表',
  `table_id` bigint NOT NULL DEFAULT '0' COMMENT '变更对应表的ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_uuid` (`user_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户余额变动表';

-- Inserting 1 row into user_accounts
-- Insert batch #1
INSERT INTO user_accounts (id, user_uuid, `before`, `change`, `after`, remark, `table`, table_id, created_at) VALUES
(1, 'TlauENMYypybGMstEfdGNPLwcDwPGg', 0, 0, 0, '', '', 0, '2022-09-01 01:43:31');

-- END TABLE user_accounts

-- BEGIN TABLE user_banks
DROP TABLE IF EXISTS user_banks;
CREATE TABLE `user_banks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mobile` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行卡手机号',
  `province` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行卡开户省',
  `city` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行卡开户市',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行卡类型',
  `abbreviation` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银联支付网关简码',
  `bank` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行名称',
  `bin_digits` int NOT NULL DEFAULT '0' COMMENT '银行卡bin码长度',
  `card_number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行卡号',
  `card_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行卡名称',
  `card_bin` int NOT NULL DEFAULT '0' COMMENT '银行卡bin码',
  `card_digits` int NOT NULL DEFAULT '0' COMMENT '银行卡号长度',
  `logo` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '银行logo',
  `is_luhn` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '是否支持luhn校验 true-校验成功 false-校验失败',
  `images` json DEFAULT NULL COMMENT '证件图片',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_uuid` (`user_uuid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户银行账户表';

-- Inserting 1 row into user_banks
-- Insert batch #1
INSERT INTO user_banks (id, user_uuid, mobile, province, city, type, abbreviation, bank, bin_digits, card_number, card_name, card_bin, card_digits, logo, is_luhn, images, created_at, updated_at, deleted_at) VALUES
(1, 'TlauENMYypybGMstEfdGNPLwcDwPGg', '', '', '', '', '', '', 0, '', '', 0, 0, '', '', NULL, '2022-09-01 01:43:47', NULL, NULL);

-- END TABLE user_banks

-- BEGIN TABLE user_groups
DROP TABLE IF EXISTS user_groups;
CREATE TABLE `user_groups` (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户组表';

-- Inserting 1 row into user_groups
-- Insert batch #1
INSERT INTO user_groups (id, group_uuid, title, content, created_at, updated_at, deleted_at) VALUES
(1, 'twaAwNRZAxnHCVHrIhCuTbyeljquHJ', 'test', '342', '2022-09-01 17:33:36', '2022-09-01 17:33:36', NULL);

-- END TABLE user_groups

-- BEGIN TABLE user_verifieds
DROP TABLE IF EXISTS user_verifieds;
CREATE TABLE `user_verifieds` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'personal, company',
  `state` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'pass, fail',
  `method` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'API认证, person本站, face人脸核身',
  `number` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证件号',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `front_photo` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '正面照, 营业执照照片',
  `back_photo` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '背面照',
  `hand_photo` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手持照片',
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地址',
  `remarks` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_uuid` (`user_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户实名认证表';

-- Inserting 1 row into user_verifieds
-- Insert batch #1
INSERT INTO user_verifieds (id, user_uuid, type, state, method, number, name, front_photo, back_photo, hand_photo, address, remarks, created_at, updated_at, deleted_at) VALUES
(1, 'TlauENMYypybGMstEfdGNPLwcDwPGg', '', '', '', '', 'test', '', '', '', '', '', '2022-09-01 08:27:15', NULL, NULL);

-- END TABLE user_verifieds

-- BEGIN TABLE users
DROP TABLE IF EXISTS users;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `wechat_openid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `nick_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `pay_password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `account_amount` decimal(15,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '账户金额',
  `groups` json DEFAULT NULL COMMENT '用户分类',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQ_USERNAME` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- Inserting 2 rows into users
-- Insert batch #1
INSERT INTO users (id, uuid, wechat_openid, name, avatar, nick_name, password, pay_password, account_amount, `groups`, created_at, updated_at, deleted_at) VALUES
(1, 'cNhVeztqAVTeTyahzxWYhcZIjtpBtK', '', 'test', '', 'Test', '$2a$14$10BXVovf/FXpbo09bw7T1ukVgkpp/pYOYKMQIPCx3b8xSJHdzfRnm', '', 0, NULL, '2022-07-13 15:01:24', '2022-09-01 01:43:08', NULL),
(3, 'TlauENMYypybGMstEfdGNPLwcDwPGg', 'oE2tk1rxzf71mx3H2NX4rNbAAqZ4', '', '', '', '', '', 0, NULL, '2022-07-16 23:52:30', '2022-09-01 01:43:08', NULL);

-- END TABLE users

