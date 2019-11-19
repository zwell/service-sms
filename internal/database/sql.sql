CREATE TABLE `supplier` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '名称',
  `code` varchar(20) NOT NULL COMMENT '唯一标识',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1可用',
  `created_at` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT '供应商';

CREATE TABLE `template` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '名称',
  `code` varchar(50) NOT NULL COMMENT '编码',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型。1验证码，2通知，3营销',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1可用',
  `created_at` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT '模板';

CREATE TABLE `template_supplier` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `supplier_id` int(11) NOT NULL COMMENT '供应商ID',
  `template_id` int(11) NOT NULL COMMENT '模板ID',
  `template_code` varchar(100) NOT NULL DEFAULT '' COMMENT '模板CODE',
  `template_params` varchar(100) NOT NULL DEFAULT '' COMMENT '模板变量',
  `template_content` varchar(100) NOT NULL DEFAULT '' COMMENT '模板内容',
  `priority` tinyint(1) NOT NULL DEFAULT 10 COMMENT '优先级',
  `price` smallint(8) NOT NULL COMMENT '价格.分',
  `created_at` int(11) NOT NULL DEFAULT '0',
  `updated_at` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  INDEX `supplier_id` (`supplier_id`) USING BTREE ,
  INDEX `template_id` (`template_id`) USING BTREE
) ENGINE=InnoDB COMMENT '供应商模板';

CREATE TABLE `send_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `task_id` int(11) NOT NULL DEFAULT 0 COMMENT '任务ID',
  `template_supplier_id` int(11) NOT NULL COMMENT '供应商模板ID',
  `phone` char(11) NOT NULL COMMENT '手机号',
  `content` varchar(100) NOT NULL COMMENT '内容',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1成功',
  `error_msg` varchar(500) NOT NULL DEFAULT '' COMMENT '错误信息',
  `created_at` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  INDEX `template_supplier_id` (`template_supplier_id`) USING BTREE
) ENGINE=InnoDB COMMENT '发送日志';