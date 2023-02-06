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
    `favourite_count` int(64) unsigned NOT NULL DEFAULT 0 COMMENT 'Favourite count',
    `comment_count`   int(64) unsigned NOT NULL DEFAULT 0 COMMENT 'Comment count',
    `title`           varchar(255) NOT NULL DEFAULT '' COMMENT 'Video title',
    `created_at`      timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Video create time',
    `updated_at`      timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Video update time',
    `deleted_at`      timestamp(3) NULL DEFAULT NULL COMMENT 'Video delete time',
    PRIMARY KEY (`id`),
    KEY               `idx_author_id` (`author_id`) COMMENT 'Author id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `relations`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `from_id`       bigint unsigned NOT NULL COMMENT 'From user id',
    `to_id`         bigint unsigned NOT NULL COMMENT 'To user id',
    `created_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Relation create time',
    `updated_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Relation update time',
    `deleted_at`    timestamp(3) NULL DEFAULT NULL COMMENT 'Relation delete time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Relation table';

CREATE TABLE `favourites`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`       bigint unsigned NOT NULL COMMENT 'User id',
    `video_id`      bigint unsigned NOT NULL COMMENT 'Video id',
    `created_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'Favourite create time',
    `updated_at`    timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'Favourite update time',
    `deleted_at`    timestamp(3) NULL DEFAULT NULL COMMENT 'Favourite delete time',
    PRIMARY KEY (`id`),
    KEY             `idx_user_id` (`user_id`) COMMENT 'User id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favourite table';
