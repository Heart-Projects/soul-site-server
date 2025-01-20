use blog;
-- 与业务无关，通用权限设计
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_identify` varchar(256) NOT NULL DEFAULT '' COMMENT '用户标识',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
    `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
    `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
    `home_url` varchar(100) DEFAULT '' COMMENT '个人主页',
    `invite_code` varchar(10) NOT NULL DEFAULT '' COMMENT '邀请码',
    `avatar` varchar(100) DEFAULT NULL COMMENT '头像',
    `background_login` tinyint(1) unsigned DEFAULT '1' COMMENT '允许后台登陆: 0- 禁用 1- 启用',
    `status` tinyint(1) unsigned DEFAULT '1' COMMENT '状态: 0- 禁用 1- 启用',
    `ip` varchar(64) NOT NULL DEFAULT '' COMMENT 'ip',
    `remark` varchar(256) NOT NULL DEFAULT '' COMMENT '备注',
    `last_login_ip` varchar(30) DEFAULT NULL COMMENT '上次登录IP',
    `last_login_at` datetime DEFAULT NULL COMMENT '上次登录时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_user_identify` (`user_identify`),
    UNIQUE KEY `uk_name` (`name`),
    UNIQUE KEY `uk_email` (`email`),
    UNIQUE KEY `uk_phone` (`phone`)
) DEFAULT CHARSET=utf8mb3 COMMENT='用户信息';

-- 设计说明： 超管拥有一切权限。
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
     `role_type` tinyint NOT NULL DEFAULT '1' COMMENT '0:超级管理员 1：非超管',
     `name` varchar(100) DEFAULT NULL COMMENT '名称',
     `description` varchar(500) DEFAULT NULL COMMENT '描述',
     `status` int DEFAULT '1' COMMENT '启用状态：0->禁用；1->启用',
     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
     UNIQUE KEY `uk_name` (`name`),
     PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb3 COMMENT='用户角色表';

DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
   `user_id` bigint  NOT NULL DEFAULT 0 COMMENT '用户id,参见sys_user',
   `role_id` bigint  NOT NULL DEFAULT 0 COMMENT '角色id,参见sys_role',
   PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb3 COMMENT='用户角色关系表';

-- 资源组的含义： 一个页面来说有url 资源/元素资源/api 接口资源。
DROP TABLE IF EXISTS `sys_resource_category`;
CREATE TABLE `sys_resource_category` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT,
     `name` varchar(200) DEFAULT NULL default '' COMMENT '分类名称',
     `is_level` tinyint(4) unsigned DEFAULT '0' COMMENT '是否具有层级: 0- 否 1- 是',
     `is_group` tinyint(4) unsigned DEFAULT '0' COMMENT '是否是资源组: 0- 否 1- 是',
     `is_simple_type` tinyint(4) unsigned DEFAULT '0' COMMENT '是否简单类型: 0- 是 1- 否',
     `schema_json` varchar(2000) not null DEFAULT '' COMMENT '自定义类型的schema',
     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
     UNIQUE KEY `uk_name` (`name`),
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='资源类别表';

-- 资源子类别表
DROP TABLE IF EXISTS `sys_resource_category_child`;
CREATE TABLE `sys_resource_category_child` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT,
     `resource_category_id` bigint unsigned NOT NULL default 0 comment '类别ID',
     `child_resource_category_id` bigint unsigned NOT NULL default 0 comment '子类别ID',
     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='资源类别子类别表';

DROP TABLE IF EXISTS `sys_resource`;
CREATE TABLE `sys_resource` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `parent_id` bigint unsigned NOT NULL default 0 comment '父资源id',
    `category_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '资源分类ID',
    `child_category_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '子资源分类ID',
    `name` varchar(200) NOT NULL DEFAULT '' COMMENT '资源名称，只用做显示用',
    `content` varchar(1000) NOT NULL DEFAULT '' COMMENT '自定义资源内容，由资源类型的json_schema决定',
    `description` varchar(500) NOT NULL DEFAULT '' COMMENT '描述',
    `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
    `create_time` datetime DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_name` (`name`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='后台资源表';

DROP TABLE IF EXISTS `sys_role_resource`;
CREATE TABLE `sys_role_resource` (
      `id` bigint unsigned  NOT NULL AUTO_INCREMENT,
      `role_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '角色ID',
      `resource_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '资源ID',
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='后台角色资源关系表';

-- 通用系统配置或者字典表
DROP TABLE IF EXISTS `system_config`;
CREATE TABLE `system_config` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
     `config_name` varchar(100)  NOT NULL DEFAULT '' COMMENT '配置编名称',
     `config_key` varchar(200)  NOT NULL DEFAULT '' COMMENT '配置健',
     `config_value` varchar(500)  NOT NULL DEFAULT '' COMMENT '配置值',
     `order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序: 0',
     `remark` varchar(255)  NOT NULL DEFAULT '' COMMENT '备注',
     `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0禁用，1-启用',
     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`),
     UNIQUE KEY `uk_config_name` (`config_name`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='通用系统配置/字典表';

-- 操作日志
DROP TABLE IF EXISTS `system_log`;
CREATE TABLE `system_log` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
     `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
     `resource_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '资源id',
     `action` int(10) unsigned DEFAULT '0' COMMENT '动作: 0- 访问 1- 新增 2- 更新 3- 删除',
     `ip` varchar(20) NOT NULL DEFAULT '',
     `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb3 COMMENT='系统记录表';
-- 系统通知
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `title` varchar(255) DEFAULT NULL COMMENT '公告标题',
      `content` varchar(5000) DEFAULT NULL COMMENT '内容',
      `status` int(4) unsigned DEFAULT '1' COMMENT '状态: 0- 待发布，1- 已发布',
      `order` int(4) DEFAULT NULL COMMENT '排序值',
      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COMMENT '系统公告表';


-- 业务系统相关
-- 业务系统相关
-- 网站主题
DROP TABLE IF EXISTS `content_topic`;
CREATE TABLE `content_topic` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
      `name` varchar(100) NOT NULL DEFAULT '' COMMENT '主题名称',
      `note` varchar(200) NOT NULL DEFAULT '' COMMENT '说明',
      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`),
      UNIQUE KEY `uk_name` (`name`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='网站话题表';

-- 全站分类
DROP TABLE IF EXISTS `site_category`;
CREATE TABLE `site_category` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `parent_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '父分类id',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '类别名称',
    `note` varchar(200) NOT NULL DEFAULT '' COMMENT '说明',
    `order` int NOT NULL DEFAULT '0' COMMENT '显示顺序：',
    `is_show` tinyint NOT NULL DEFAULT '0' COMMENT '是否显示：0- 不显示 1- 显示',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`name`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='站点类别表';

-- 专栏主题表
DROP TABLE IF EXISTS `article_column_theme`;
CREATE TABLE `article_column_theme` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '主题名称',
  `note` varchar(200) NOT NULL DEFAULT '' COMMENT '说明',
  `icon` varchar(500) NOT NULL DEFAULT '' COMMENT '图片链接',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='专栏主题表';

-- 专栏类别
DROP TABLE IF EXISTS `article_column_category`;
CREATE TABLE `article_column_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '父分类id',
  `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '专栏类别名称',
  `level` int(10) unsigned NOT NULL DEFAULT 1 COMMENT '级别',
  `tree_path` varchar(200) NOT NULL DEFAULT '/' COMMENT '层次路径',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='专栏类别表';

-- 文章专栏
DROP TABLE IF EXISTS `article_column`;
CREATE TABLE `article_column` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
      `category_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '类别ID',
      `theme_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '主题id',
      `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
      `identify` varchar(100) NOT NULL DEFAULT '' COMMENT '专栏唯一标识',
      `name` varchar(100) NOT NULL DEFAULT '' COMMENT '专栏名称',
      `note` varchar(200) NOT NULL DEFAULT '' COMMENT '说明',
      `summary` varchar(200) NOT NULL DEFAULT '' COMMENT '简介',
      `icon` varchar(500) NOT NULL DEFAULT '' COMMENT '图片链接',
      `view_scope` tinyint NOT NULL DEFAULT '0' COMMENT '查看权限: 0-私有，1-公开',
      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`),
      UNIQUE KEY `uk_name` (`name`) USING BTREE,
      UNIQUE KEY `uk_identify` (`identify`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='专栏表';

-- 专栏文章
DROP TABLE IF EXISTS `column_articles`;
CREATE TABLE `column_articles` (
   `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章ID',
   `parent_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '父id',
   `column_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '专栏id',
   `type` varchar(50) NOT NULL DEFAULT 'doc' COMMENT '类型: 0- doc, 1-folder',
   `identify` varchar(100) NOT NULL DEFAULT '' COMMENT '文章唯一标识',
   `title` varchar(255) DEFAULT NULL COMMENT '标题/文章分类名称',
   `sort_number` int(11) NOT NULL DEFAULT 0 COMMENT '排序值',
   `tree_path` varchar(500) NOT NULL DEFAULT '/' COMMENT '专栏层次路径',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_c_identify` ( `column_id`,`identify`)
) DEFAULT CHARSET=utf8 COMMENT='专栏文章表';

-- 文章内容表，专门存储内容
DROP TABLE IF EXISTS `article_content`;
CREATE TABLE `article_content` (
      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `source` int NOT NULL DEFAULT '0' COMMENT '类型: 0- 普通文章, 1- 专栏',
      `doc_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '文档id',
      `content` mediumtext COMMENT '内容',
      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
      PRIMARY KEY (`id`),
      unique key `uk_doc_id` (`source`,`doc_id`)
) DEFAULT CHARSET=utf8 COMMENT='文章内容表';

-- 文章分类表
DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
   `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
   `parent_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '父分类id',
   `name` varchar(100) NOT NULL DEFAULT '' COMMENT '类别名称',
   `note` varchar(200) NOT NULL DEFAULT '' COMMENT '说明',
   `order` int NOT NULL DEFAULT '0' COMMENT '显示顺序：',
   `level` int(10) unsigned NOT NULL DEFAULT 1 COMMENT '级别',
   `is_show` tinyint NOT NULL DEFAULT '0' COMMENT '是否显示：0- 不显示 1- 显示',
   `article_count` int NOT NULL DEFAULT '0' COMMENT '文章数目',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`),
   UNIQUE KEY `uk_name` (`user_id`, `name`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='文章类别表';
-- 标签
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '标签名称',
    `note` varchar(200) NOT NULL DEFAULT '' COMMENT '说明',
    `color` varchar(200) NOT NULL DEFAULT '' COMMENT '颜色',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`user_id`, `name`) USING BTREE
) DEFAULT CHARSET=utf8mb3 COMMENT='文章类别表';

-- 文章
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
   `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章ID',
   `user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
   `site_category_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '站点类别',
   `category_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '所属类别',
   `column_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '专栏id',
   `type` int NOT NULL DEFAULT '1' COMMENT '类型: 0-文章分类 1-普通文章，2-专栏文章',
   `category_tree_path` varchar(500) NOT NULL DEFAULT '/' COMMENT '类别层次路径',
   `column_tree_path` varchar(500) NOT NULL DEFAULT '/' COMMENT '专栏层次路径',
   `title` varchar(255) DEFAULT NULL COMMENT '标题/文章分类名称',
   `summary` varchar(500) COMMENT '摘要',
   `thumbnail` varchar(500) DEFAULT NULL COMMENT '缩略图',
   `content` mediumtext COMMENT '内容',
   `view_count` int(11) DEFAULT '0' COMMENT '访问量',
   `comment_count` int(11) DEFAULT '0' COMMENT '评论数',
   `like_count` int(11) DEFAULT '0' COMMENT '点赞数',
   `is_comment` int(4) unsigned DEFAULT NULL COMMENT '是否允许评论',
   `is_on_top` int(4) unsigned DEFAULT '1' COMMENT '是否置顶',
   `status` int(4) unsigned DEFAULT '1' COMMENT '状态： 0- 草稿 1-发布 2- 垃圾箱',
   `order` int(11) unsigned DEFAULT NULL COMMENT '排序值',
   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COMMENT='文章表';

-- 文章标签关联表
DROP TABLE IF EXISTS `article_tag_relation`;
CREATE TABLE `article_tag_relation` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章ID',
    `article_id` int(11) NOT NULL COMMENT '文章ID',
    `tag_id` int(11) NOT NULL COMMENT '标签ID',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 文章评论表
DROP TABLE IF EXISTS `article_comment`;
CREATE TABLE `article_comment` (
       `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
       `parent_id` bigint unsigned NOT NULL  DEFAULT '0' COMMENT '上级评论ID',
       `article_id` bigint unsigned NOT NULL  DEFAULT '0' COMMENT '文章ID',
       `user_id` bigint unsigned NOT NULL  DEFAULT '0' COMMENT '评论人ID',
       `user_nick_name` varchar(50) DEFAULT NULL COMMENT '评论人昵称',
       `user_home_url` varchar(50) DEFAULT NULL COMMENT '评论人个人主页',
       `user_avatar` varchar(100) DEFAULT NULL COMMENT '评论人头像',
       `comment_content` varchar(1000) DEFAULT NULL COMMENT '内容',
       `agent` varchar(200) DEFAULT NULL COMMENT '浏览器信息',
       `ip` varchar(50) DEFAULT NULL COMMENT 'IP',
       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
       PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT '文章评论表';

CREATE TABLE `article_statistics_data` (
       `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
       `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
       `article_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
       `view_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '浏览数',
       `comment_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
       `like_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
       `favorite_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
       `forward_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '转发数',
       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
       `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
       PRIMARY KEY (`id`),
       UNIQUE KEY `uk_name` (`user_id`,`article_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='文章数据统计表';

CREATE TABLE `user_article_statistics_data` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `hot_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '浏览数',
    `article_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '文章数',
    `follow_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '关注数',
    `favorite_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
    `forward_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '转发数',
    `comment_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_name` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户文章数据统计表'








