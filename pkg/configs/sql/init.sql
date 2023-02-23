CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `follow_count` bigint NOT NULL DEFAULT 0 comment '关注总数',
    `follower_count` bigint NOT NULL DEFAULT 0 comment '粉丝总数',
    `video_count` bigint NOT NULL DEFAULT 0 comment '视频总数',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_name` (`user_name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `likes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint(20) NOT NULL COMMENT '点赞用户id',
  `video_id` bigint(20) NOT NULL COMMENT '被点赞的视频id',
  `cancel` tinyint(4) NOT NULL DEFAULT '0' COMMENT '默认点赞为0，取消赞为1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userIdtoVideoIdIdx` (`user_id`,`video_id`) USING BTREE,
  KEY `userIdIdx` (`user_id`) USING BTREE,
  KEY `videoIdx` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1229 DEFAULT CHARSET=utf8 COMMENT='点赞表';
