/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost:3306
 Source Schema         : blog_data

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 06/04/2022 13:39:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_article
-- ----------------------------
DROP TABLE IF EXISTS `t_article`;
CREATE TABLE `t_article` (
  `article_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `title` varchar(150) NOT NULL DEFAULT '' COMMENT '文章标题',
  `description` varchar(1000) NOT NULL DEFAULT '' COMMENT '文章描述',
  `content` longtext NOT NULL COMMENT '文章内容',
  `cover_img_url` varchar(200) NOT NULL DEFAULT '' COMMENT '文章封面图',
  `visits` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章浏览次数',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '文章排序',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '文章状态（0：禁用,1：启用）',
  `is_deleted` tinyint(3) NOT NULL COMMENT '是否删除（0：正常,1：已删除）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`article_id`) USING BTREE,
  KEY `idx_status` (`status`),
  KEY `idx_is_deleted` (`is_deleted`),
  KEY `idx_sort` (`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='文章信息';

-- ----------------------------
-- Records of t_article
-- ----------------------------
BEGIN;
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'this is a test', '', 'xxxxxxxx', '', 0, 0, 0, 0, '2021-02-12 21:43:21', '2021-02-12 21:43:21', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'test', '', 'xxxxxxxx', '', 0, 0, 0, 0, '2021-02-12 21:43:53', '2021-02-12 21:43:53', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'asdasd', '', 'xxxxxxxx', '', 0, 0, 0, 0, '2021-02-12 21:43:58', '2021-02-12 21:43:58', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'asfdasfasf', '', 'xxxxxxxx', '', 0, 0, 0, 0, '2021-02-12 21:44:00', '2021-02-12 21:44:00', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '34234234', '', 'xxxxxxxx', '', 0, 0, 0, 0, '2021-02-12 21:44:02', '2021-02-12 21:44:02', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'asdasdasd', 'ahdjkhajkhdjkahsjkdhasjkdhjkhjkhdajkshdjkahsjkdhajkshdjkhasjkdhjkasdasd', 'xxxxxxxx', '', 0, 0, 0, 0, '2021-02-12 21:44:38', '2021-02-12 21:44:38', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'sdfs', 'ahdjkhajkhdjkahsjkdhasjkdhjkhjkhdajkshdjkahsjkdhajkshdjkhasjkdhjkasdasd', 'xxxxxxxx', 'shhshs', 0, 0, 0, 0, '2021-02-12 21:45:15', '2021-02-12 21:45:15', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 'dasdas', 'ahdjkhajkhdjkahsjkdhasjkdhjkhjkhdajkshdjkahsjkdhajkshdjkhasjkdhjkasdasd', 'xxxxxxxx', 'shhshs', 0, 100, 1, 0, '2021-02-12 21:45:40', '2021-02-12 21:45:40', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'asfasdfas', 'ahdjkhajkhdjkahsjkdhasjkdhjkhjkhdajkshdjkahsjkdhajkshdjkhasjkdhjkasdasd', 'xxxxxxxx', 'shhshs', 0, 100, 1, 0, '2021-02-12 21:45:42', '2021-02-12 21:45:42', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 'asfasfasf', 'ahdjkhajkhdjkahsjkdhasjkdhjkhjkhdajkshdjkahsjkdhajkshdjkhasjkdhjkasdasd', 'xxxxxxxx', 'shhshs', 0, 100, 1, 0, '2021-02-12 21:45:44', '2021-02-12 21:45:44', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 'asdasfa', 'ahdjkhajkhdjkahsjkdhasjkdhjkhjkhdajkshdjkahsjkdhajkshdjkhasjkdhjkasdasd', 'xxxxxxxx', 'shhshs', 0, 100, 1, 0, '2021-02-12 21:46:05', '2021-02-12 21:46:05', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 'asdfasf', 'ahdjkhajkhdjkahsjkdhasjkdhjkhjkhdajkshdjkahsjkdhajkshdjkhasjkdhjkasdasd', 'xxxxxxxx', 'shhshs', 0, 100, 1, 0, '2021-02-12 21:46:07', '2021-02-12 21:46:07', NULL);
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 'afasfaf', '123', '123', '123', 0, 123, 0, 0, '2021-02-12 21:46:14', '2021-02-12 21:57:20', '2021-02-12 22:02:08');
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 'asd', '123', '123', '123', 0, 123, 0, 0, '2021-02-12 21:46:19', '2021-02-12 21:57:11', '2021-02-12 22:02:27');
INSERT INTO `t_article` (`article_id`, `title`, `description`, `content`, `cover_img_url`, `visits`, `sort`, `status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 'tests', '123', '123', '123', 0, 123, 0, 0, '2021-02-12 21:46:22', '2021-02-12 21:56:43', '2021-02-12 22:02:27');
COMMIT;

-- ----------------------------
-- Table structure for t_article_category
-- ----------------------------
DROP TABLE IF EXISTS `t_article_category`;
CREATE TABLE `t_article_category` (
  `category_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章分类ID',
  `category_pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级分类ID',
  `category_name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
  `category_sort` int(11) NOT NULL DEFAULT '0' COMMENT '分类排序',
  `is_deleted` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除（0：正常，1：已删除）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章分类信息';

-- ----------------------------
-- Records of t_article_category
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_article_tag`;
CREATE TABLE `t_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号',
  `article_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章编号',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签编号',
  `is_deleted` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除（0：正常,1：已删除）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_tag_id` (`tag_id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_is_deleted` (`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联表';

-- ----------------------------
-- Records of t_article_tag
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for t_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_tag`;
CREATE TABLE `t_tag` (
  `tag_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `tag_name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名称',
  `tag_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '标签状态（0：禁用,1：启用）',
  `is_deleted` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除（0：正常,1：已删除）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`tag_id`) USING BTREE,
  KEY `idx_tag_status` (`tag_status`),
  KEY `idx_tag_deleted` (`is_deleted`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='标签信息';

-- ----------------------------
-- Records of t_tag
-- ----------------------------
BEGIN;
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '', 0, 0, '2021-01-29 21:12:21', '2021-01-29 21:12:21', '2021-02-12 20:26:03');
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '', 0, 0, '2021-01-29 21:30:08', '2021-01-29 21:30:08', '2021-02-12 20:26:03');
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'golang', 1, 0, '2021-02-12 20:03:56', '2021-02-12 20:03:56', '2021-02-12 20:27:46');
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'php', 1, 0, '2021-02-12 20:04:34', '2021-02-12 20:04:34', '2021-02-12 20:25:52');
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'redis', 1, 0, '2021-02-12 20:04:40', '2021-02-12 20:04:40', '2021-02-12 20:27:46');
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'mySQL2777', 0, 0, '2021-02-12 20:04:45', '2021-02-12 20:14:32', '2021-02-12 20:25:17');
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'mySQL', 1, 0, '2021-02-12 20:32:43', '2021-02-12 20:32:43', NULL);
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 'docker', 1, 0, '2021-02-12 20:45:23', '2021-02-12 20:45:23', NULL);
INSERT INTO `t_tag` (`tag_id`, `tag_name`, `tag_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'k8s', 1, 0, '2021-02-12 20:45:28', '2021-02-12 20:45:28', NULL);
COMMIT;

-- ----------------------------
-- Table structure for t_website
-- ----------------------------
DROP TABLE IF EXISTS `t_website`;
CREATE TABLE `t_website` (
  `web_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '网站ID',
  `web_name` varchar(150) NOT NULL DEFAULT '' COMMENT '网站名称',
  `web_url` varchar(200) NOT NULL DEFAULT '' COMMENT '网站地址',
  `web_description` varchar(1000) NOT NULL DEFAULT '' COMMENT '网站描述',
  `web_sort` int(11) NOT NULL DEFAULT '0' COMMENT '列表排序',
  `web_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '网站状态（0：不显示,1：显示）',
  `is_deleted` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除（0：正常,1：已删除）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`web_id`) USING BTREE,
  UNIQUE KEY `unique_web_name` (`web_name`),
  UNIQUE KEY `unique_web_url` (`web_url`),
  KEY `idx_web_sort` (`web_sort`),
  KEY `idx_web_status` (`web_status`),
  KEY `idx_is_deleted` (`is_deleted`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='网站信息';

-- ----------------------------
-- Records of t_website
-- ----------------------------
BEGIN;
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '掘金', 'https://juejin.cn/', '掘金 - 代码不止，掘金不停', 10, 0, 0, '2021-12-09 16:43:14', '2021-12-13 10:39:41', '2021-12-13 10:46:22');
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'GitHub', 'https://github.com/', 'GitHub', 100, 0, 0, '2021-12-09 17:11:29', '2021-12-09 17:11:29', '2021-12-13 10:46:22');
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'StackOverflow', 'https://stackoverflow.com/', 'Stack Overflow - Where Developers Learn, Share, & Build Careers', 80, 0, 0, '2021-12-09 17:14:50', '2021-12-09 17:14:50', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'CSDN', 'https://www.csdn.net/', 'CSDN - 专业开发者社区', 0, 0, 0, '2021-12-09 17:17:38', '2021-12-09 17:17:38', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '简书', 'https://www.jianshu.com/', '简书 - 创作你的创作', 0, 0, 0, '2021-12-09 17:19:29', '2021-12-09 17:19:29', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '知乎', 'https://www.zhihu.com/', '知乎', 0, 0, 0, '2021-12-09 17:22:27', '2021-12-09 17:22:27', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'InfoQ', 'https://www.infoq.cn/', 'InfoQ - 促进软件开发及相关领域知识与创新的传播-极客邦', 0, 0, 0, '2021-12-09 17:23:19', '2021-12-09 17:23:19', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 'SegmentFault', 'https://segmentfault.com/', 'SegmentFault 思否', 40, 0, 0, '2021-12-09 17:24:30', '2021-12-09 17:24:30', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'cnblogs', 'https://www.cnblogs.com/', '博客园 - 开发者的网上家园', 0, 0, 0, '2021-12-09 17:25:15', '2021-12-09 17:25:15', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 'OSCHINA', 'https://www.oschina.net', 'OSCHINA - 中文开源技术交流社区', 0, 0, 0, '2021-12-09 17:29:15', '2021-12-09 17:29:15', NULL);
INSERT INTO `t_website` (`web_id`, `web_name`, `web_url`, `web_description`, `web_sort`, `web_status`, `is_deleted`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, '百度搜索test', 'https://www.baidu.com/', '百度搜索,百度搜索', 201, 1, 1, '2021-12-10 21:47:11', '2021-12-13 10:28:46', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
