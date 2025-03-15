CREATE TABLE `sequence` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `stub` VARCHAR(1) NOT NULL COMMENT 'place hold',
    `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'timestamp',
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_idx_stub` (`stub`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='long-to-short url map table';