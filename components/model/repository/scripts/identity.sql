USE `###DATABASE###_history`;

CREATE TABLE `identities`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `username` VARCHAR(32) NOT NULL,
    `phone_number` VARCHAR(12) NOT NULL,
    `phone_number_confirmed` BIT(1) NOT NULL,
    `first_name` VARCHAR(128) NOT NULL,
    `last_name` VARCHAR(128) NOT NULL,
    `display_name` VARCHAR(128) NOT NULL,
    `email` VARCHAR(128) NOT NULL,
    `email_confirmed` BIT(1) NOT NULL,
    `avatar` VARCHAR(512) NOT NULL,
    `banner` VARCHAR(512) NOT NULL,
    `summary` VARCHAR(512) NOT NULL,
    `token` VARCHAR(256) NOT NULL,
    `multi_factor` BIT(1) NOT NULL,
    `hash` VARCHAR(256) NOT NULL,
    `salt` VARCHAR(64) NOT NULL,
    `public_key` VARCHAR(4096) NOT NULL,
    `private_key` VARCHAR(4096) NOT NULL,
    `permission` BIGINT UNSIGNED NOT NULL,
    `restriction` INT UNSIGNED NOT NULL,
    `last_login` BIGINT NOT NULL,
    `login_count` INT UNSIGNED NOT NULL,
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

CREATE TABLE `identities`
(
    `id` BIGINT NOT NULL,
    `username` VARCHAR(32) NOT NULL UNIQUE,
    `phone_number` VARCHAR(12) NOT NULL,
    `phone_number_confirmed` BIT(1) NOT NULL,
    `first_name` VARCHAR(128) NOT NULL,
    `last_name` VARCHAR(128) NOT NULL,
    `display_name` VARCHAR(128) NOT NULL,
    `email` VARCHAR(128) NOT NULL UNIQUE,
    `email_confirmed` BIT(1) NOT NULL,
    `avatar` VARCHAR(512) NOT NULL,
    `banner` VARCHAR(512) NOT NULL,
    `summary` VARCHAR(512) NOT NULL,
    `token` VARCHAR(256) NOT NULL UNIQUE,
    `multi_factor` BIT(1) NOT NULL,
    `hash` VARCHAR(256) NOT NULL,
    `salt` VARCHAR(64) NOT NULL,
    `public_key` VARCHAR(4096) NOT NULL,
    `private_key` VARCHAR(4096) NOT NULL,
    `permission` BIGINT UNSIGNED NOT NULL,
    `restriction` INT UNSIGNED NOT NULL,
    `last_login` BIGINT NOT NULL,
    `login_count` INT UNSIGNED NOT NULL,
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
