CREATE TABLE `users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'User account create time',
    `updated_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'User account update time',
    `deleted_at` timestamp(3) NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `videos`
(
    `id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `author_id`       bigint unsigned NOT NULL COMMENT 'Author id',
    `play_url`        varchar(255) NOT NULL DEFAULT '' COMMENT 'Play video url',
    `cover_url`       varchar(255) NOT NULL DEFAULT '' COMMENT 'Video cover url',
    `title`           varchar(255) NOT NULL DEFAULT '' COMMENT 'Video title',
    `created_at`      timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Video create time',
    `updated_at`      timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Video update time',
    `deleted_at`      timestamp(3) NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY               `idx_author_id` (`author_id`) COMMENT 'Author id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `follows`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_id`       bigint unsigned NOT NULL COMMENT 'From user id',
    `to_id`         bigint unsigned NOT NULL COMMENT 'To user id',
    `created_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'follow create time',
    `updated_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'follow update time',
    `deleted_at`    timestamp(3) NULL DEFAULT NULL COMMENT 'follow delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_from_id` (`from_id`) COMMENT 'from_id index',
    KEY          `idx_to_id` (`to_id`) COMMENT 'to_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Relation table';

CREATE TABLE `messages`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`       bigint unsigned NOT NULL COMMENT 'From user id',
    `to_user_id`         bigint unsigned NOT NULL COMMENT 'To user id',
    `content`       text NOT NULL COMMENT 'content context',
    `created_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'message create time',
    `updated_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'message update time',
    `deleted_at`    timestamp(3) NULL DEFAULT NULL COMMENT 'message delete time',
    PRIMARY KEY (`id`),
    KEY             `idx_user_id` (`user_id`) COMMENT 'user_id index',
    KEY             `idx_to_user_id` (`to_user_id`) COMMENT 'to_user_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='message table';

CREATE TABLE `favorites`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`       bigint unsigned NOT NULL COMMENT 'User id',
    `video_id`      bigint unsigned NOT NULL COMMENT 'Video id',
    `created_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Favourite create time',
    `updated_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Favourite update time',
    `deleted_at`    timestamp(3) NULL DEFAULT NULL COMMENT 'Favourite delete time',
    PRIMARY KEY (`id`),
    KEY             `idx_user_video_id` (`user_id`, `video_id`) COMMENT 'User-Video id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favourite table';

CREATE TABLE `comments`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`       bigint unsigned NOT NULL COMMENT 'User id',
    `video_id`      bigint unsigned NOT NULL COMMENT 'Video id',
    `content`       text NOT NULL COMMENT 'Comment context',
    `create_date`   varchar(16) NOT NULL COMMENT 'Comment create date',
    `created_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Comment create time',
    `updated_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Comment update time',
    `deleted_at`    timestamp(3) NULL DEFAULT NULL COMMENT 'Comment delete time',
    PRIMARY KEY (`id`),
    KEY             `idx_video_id` (`video_id`) COMMENT 'Video id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';

