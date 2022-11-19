-- 创建数据库
CREATE DATABASE IF NOT EXISTS `x_db`;
USE `x_db`;
CREATE TABLE `medal_info_tab` (
                                  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                  `user_did` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
                                  `event_uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动uid',
                                  `teacher` varchar(16) NOT NULL DEFAULT '',
                                  `medal_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '徽章状态',
                                  `chain_hash` varchar(128) NOT NULL DEFAULT '',
                                  `chain_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '上链状态 0:未上链 1:队列中 2:已上链 3:上链失败',
                                  `ctime` int(10) NOT NULL DEFAULT '0',
                                  `mtime` int(10) NOT NULL DEFAULT '0',
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='徽章表';