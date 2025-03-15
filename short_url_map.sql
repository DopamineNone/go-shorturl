CREATE TABLE `short_url_map` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `is_del` TINYINT NOT NULL DEFAULT '0',

    `lurl` VARCHAR(160) NOT NULL,
    `md5` CHAR(32) NOT NULL,
    `surl` VARCHAR(11) NOT NULL,

    PRIMARY KEY (`id`),
    INDEX(`is_del`),
    UNIQUE(`md5`),
    UNIQUE(`surl`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = 'short url map';