CREATE TABLE `users`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`       varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
    `password`       varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `follow_count`   int unsigned NOT NULL DEFAULT 0 COMMENT 'User follow count',
    `follower_count` int unsigned NOT NULL DEFAULT 0 COMMENT 'User follower count',
    `created_at`     timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'User account create time',
    `updated_at`     timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP (3) COMMENT 'User account update time',
    `deleted_at`     timestamp(3) NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_username` (`username`) COMMENT 'Unique username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `videos`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `author_id`      bigint unsigned NOT NULL COMMENT 'Author id',
    `play_url`       varchar(255) NOT NULL DEFAULT '' COMMENT 'Play video url',
    `cover_url`      varchar(255) NOT NULL DEFAULT '' COMMENT 'Video cover url',
    `title`          varchar(255) NOT NULL DEFAULT '' COMMENT 'Video title',
    `favorite_count` int unsigned NOT NULL DEFAULT 0 COMMENT 'Video favorite count',
    `comment_count`  int unsigned NOT NULL DEFAULT 0 COMMENT 'Video Favorite count',
    `created_at`     timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Video create time',
    `updated_at`     timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP (3) COMMENT 'Video update time',
    `deleted_at`     timestamp(3) NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY              `idx_author_id` (`author_id`) COMMENT 'Author id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `follows`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_id`    bigint unsigned NOT NULL COMMENT 'From user id',
    `to_id`      bigint unsigned NOT NULL COMMENT 'To user id',
    `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Follow create time',
    `updated_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP (3) COMMENT 'Follow update time',
    `deleted_at` timestamp(3) NULL DEFAULT NULL COMMENT 'Follow delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_follow` (`from_id`, `to_id`, `deleted_at`) COMMENT 'Unique compound follow index',
    KEY          `idx_from_id` (`from_id`) COMMENT 'From user id index',
    KEY          `idx_to_id` (`to_id`) COMMENT 'To user id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Follow table';

CREATE TABLE `messages`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    bigint unsigned NOT NULL COMMENT 'From user id',
    `to_user_id` bigint unsigned NOT NULL COMMENT 'To user id',
    `content`    text         NOT NULL COMMENT 'Message content',
    `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Message create time',
    `updated_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP (3) COMMENT 'Message update time',
    `deleted_at` timestamp(3) NULL DEFAULT NULL COMMENT 'Message delete time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Message table';

CREATE TABLE `favorites`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`    bigint unsigned NOT NULL COMMENT 'User id',
    `video_id`   bigint unsigned NOT NULL COMMENT 'Video id',
    `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Favourite create time',
    `updated_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP (3) COMMENT 'Favourite update time',
    `deleted_at` timestamp(3) NULL DEFAULT NULL COMMENT 'Favourite delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY          `idx_user_video_id` (`user_id`, `video_id`, `deleted_at`) COMMENT 'User-Video id index',
    KEY `idx_user_id` (`user_id`) COMMENT 'User id index',
    KEY `idex_video_id` (`video_id`) COMMENT 'Video id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favourite table';

CREATE TABLE `comments`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`     bigint unsigned NOT NULL COMMENT 'User id',
    `video_id`    bigint unsigned NOT NULL COMMENT 'Video id',
    `content`     text         NOT NULL COMMENT 'Comment context',
    `create_date` varchar(16)  NOT NULL COMMENT 'Comment create date',
    `created_at`  timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Comment create time',
    `updated_at`  timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP (3) COMMENT 'Comment update time',
    `deleted_at`  timestamp(3) NULL DEFAULT NULL COMMENT 'Comment delete time',
    PRIMARY KEY (`id`),
    KEY           `idx_video_id` (`video_id`) COMMENT 'Video id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';

