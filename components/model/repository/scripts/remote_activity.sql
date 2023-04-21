USE `###DATABASE###_history`;

CREATE TABLE `remote_activities`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `entry_point` VARCHAR(256) NOT NULL,
    `duration` BIGINT NOT NULL,
    `successful` BIT(1) NOT NULL,
    `error_message` VARCHAR(1024) NOT NULL,
    `remote_address` VARCHAR(128) NOT NULL,
    `user_agent` VARCHAR(512) NOT NULL,
    `event_type` INT UNSIGNED NOT NULL,
    `timestamp` BIGINT NOT NULL,
    `editor` BIGINT NOT NULL,
    `status` BIGINT NOT NULL,
    `sort_order` FLOAT NOT NULL,
    `queued_at` BIGINT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `triggered_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `changed_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `payload` JSON NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  COLLATE = `utf8mb4_unicode_ci`;

USE `###DATABASE###`;

CREATE TABLE `remote_activities`
(
    `id` BIGINT NOT NULL,
    `entry_point` VARCHAR(256) NOT NULL,
    `duration` BIGINT NOT NULL,
    `successful` BIT(1) NOT NULL,
    `error_message` VARCHAR(1024) NOT NULL,
    `remote_address` VARCHAR(128) NOT NULL,
    `user_agent` VARCHAR(512) NOT NULL,
    `event_type` INT UNSIGNED NOT NULL,
    `timestamp` BIGINT NOT NULL,
    `editor` BIGINT NOT NULL DEFAULT 0,
    `status` BIGINT NOT NULL DEFAULT 0,
    `sort_order` FLOAT NOT NULL DEFAULT 0,
    `queued_at` BIGINT NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `payload` JSON NULL,
    PRIMARY KEY (`id`),
    INDEX (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  COLLATE = `utf8mb4_unicode_ci`;
