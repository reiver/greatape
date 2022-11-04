DROP DATABASE IF EXISTS `greatape_dev_history`;
CREATE DATABASE `greatape_dev_history` CHARSET = `utf8mb4` COLLATE = `utf8mb4_unicode_ci`;
USE `greatape_dev_history`;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Documents
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `documents`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `content` JSON NOT NULL,
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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	SystemSchedules
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `system_schedules`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `enabled` BIT(1) NOT NULL,
    `config` VARCHAR(1024) NOT NULL,
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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Identities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	AccessControls
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `access_controls`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `key` BIGINT UNSIGNED NOT NULL,
    `value` BIGINT UNSIGNED NOT NULL,
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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	RemoteActivities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	CategoryTypes
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `category_types`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `description` VARCHAR(64) NOT NULL,
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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Categories
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `categories`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `category_type_id` BIGINT NOT NULL,
    `category_id` BIGINT NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `description` VARCHAR(64) NOT NULL,
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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Users
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `users`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `github` VARCHAR(512) NOT NULL,
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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	ActivityPubIncomingActivities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `activity_pub_incoming_activities`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `identity_id` BIGINT NOT NULL,
    `unique_identifier` VARCHAR(128) NOT NULL,
    `timestamp` BIGINT NOT NULL,
    `from` VARCHAR(256) NOT NULL,
    `to` VARCHAR(256) NOT NULL,
    `content` VARCHAR(4096) NOT NULL,
    `raw` JSON NOT NULL,
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

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	ActivityPubOutgoingActivities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `activity_pub_outgoing_activities`
(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `action` VARCHAR(16) NOT NULL,
    `original_id` BIGINT NOT NULL,
    `identity_id` BIGINT NOT NULL,
    `unique_identifier` VARCHAR(128) NOT NULL,
    `timestamp` BIGINT NOT NULL,
    `from` VARCHAR(256) NOT NULL,
    `to` VARCHAR(256) NOT NULL,
    `content` VARCHAR(4096) NOT NULL,
    `raw` JSON NOT NULL,
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

# ══════════════════════════════════════════════════════════════════════════════════════════════════════════════

DROP DATABASE IF EXISTS `greatape_dev`;
CREATE DATABASE `greatape_dev` CHARSET = `utf8mb4` COLLATE = `utf8mb4_unicode_ci`;
USE `greatape_dev`;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Documents
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `documents`
(
    `id` BIGINT NOT NULL,
    `content` JSON NOT NULL,
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

DELIMITER $$

CREATE TRIGGER `documents_after_update`
AFTER UPDATE
ON `documents` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`documents`(`action`, `original_id`, `content`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`content`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `documents_after_delete`
AFTER DELETE
ON `documents` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`documents`(`action`, `original_id`, `content`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`content`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	SystemSchedules
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `system_schedules`
(
    `id` BIGINT NOT NULL,
    `enabled` BIT(1) NOT NULL,
    `config` VARCHAR(1024) NOT NULL,
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

DELIMITER $$

CREATE TRIGGER `system_schedules_after_update`
AFTER UPDATE
ON `system_schedules` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`system_schedules`(`action`, `original_id`, `enabled`, `config`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`enabled`, `old`.`config`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `system_schedules_after_delete`
AFTER DELETE
ON `system_schedules` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`system_schedules`(`action`, `original_id`, `enabled`, `config`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`enabled`, `old`.`config`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Identities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

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

DELIMITER $$

CREATE TRIGGER `identities_after_update`
AFTER UPDATE
ON `identities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`identities`(`action`, `original_id`, `username`, `phone_number`, `phone_number_confirmed`, `first_name`, `last_name`, `display_name`, `email`, `email_confirmed`, `avatar`, `banner`, `summary`, `token`, `multi_factor`, `hash`, `salt`, `public_key`, `private_key`, `permission`, `restriction`, `last_login`, `login_count`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`username`, `old`.`phone_number`, `old`.`phone_number_confirmed`, `old`.`first_name`, `old`.`last_name`, `old`.`display_name`, `old`.`email`, `old`.`email_confirmed`, `old`.`avatar`, `old`.`banner`, `old`.`summary`, `old`.`token`, `old`.`multi_factor`, `old`.`hash`, `old`.`salt`, `old`.`public_key`, `old`.`private_key`, `old`.`permission`, `old`.`restriction`, `old`.`last_login`, `old`.`login_count`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `identities_after_delete`
AFTER DELETE
ON `identities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`identities`(`action`, `original_id`, `username`, `phone_number`, `phone_number_confirmed`, `first_name`, `last_name`, `display_name`, `email`, `email_confirmed`, `avatar`, `banner`, `summary`, `token`, `multi_factor`, `hash`, `salt`, `public_key`, `private_key`, `permission`, `restriction`, `last_login`, `login_count`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`username`, `old`.`phone_number`, `old`.`phone_number_confirmed`, `old`.`first_name`, `old`.`last_name`, `old`.`display_name`, `old`.`email`, `old`.`email_confirmed`, `old`.`avatar`, `old`.`banner`, `old`.`summary`, `old`.`token`, `old`.`multi_factor`, `old`.`hash`, `old`.`salt`, `old`.`public_key`, `old`.`private_key`, `old`.`permission`, `old`.`restriction`, `old`.`last_login`, `old`.`login_count`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	AccessControls
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `access_controls`
(
    `id` BIGINT NOT NULL,
    `key` BIGINT UNSIGNED NOT NULL,
    `value` BIGINT UNSIGNED NOT NULL,
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

DELIMITER $$

CREATE TRIGGER `access_controls_after_update`
AFTER UPDATE
ON `access_controls` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`access_controls`(`action`, `original_id`, `key`, `value`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`key`, `old`.`value`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `access_controls_after_delete`
AFTER DELETE
ON `access_controls` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`access_controls`(`action`, `original_id`, `key`, `value`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`key`, `old`.`value`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	RemoteActivities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

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

DELIMITER $$

CREATE TRIGGER `remote_activities_after_update`
AFTER UPDATE
ON `remote_activities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`remote_activities`(`action`, `original_id`, `entry_point`, `duration`, `successful`, `error_message`, `remote_address`, `user_agent`, `event_type`, `timestamp`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`entry_point`, `old`.`duration`, `old`.`successful`, `old`.`error_message`, `old`.`remote_address`, `old`.`user_agent`, `old`.`event_type`, `old`.`timestamp`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `remote_activities_after_delete`
AFTER DELETE
ON `remote_activities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`remote_activities`(`action`, `original_id`, `entry_point`, `duration`, `successful`, `error_message`, `remote_address`, `user_agent`, `event_type`, `timestamp`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`entry_point`, `old`.`duration`, `old`.`successful`, `old`.`error_message`, `old`.`remote_address`, `old`.`user_agent`, `old`.`event_type`, `old`.`timestamp`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	CategoryTypes
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `category_types`
(
    `id` BIGINT NOT NULL,
    `description` VARCHAR(64) NOT NULL,
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

DELIMITER $$

CREATE TRIGGER `category_types_after_update`
AFTER UPDATE
ON `category_types` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`category_types`(`action`, `original_id`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `category_types_after_delete`
AFTER DELETE
ON `category_types` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`category_types`(`action`, `original_id`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Categories
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `categories`
(
    `id` BIGINT NOT NULL,
    `category_type_id` BIGINT NOT NULL,
    `category_id` BIGINT NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `description` VARCHAR(64) NOT NULL,
    `editor` BIGINT NOT NULL DEFAULT 0,
    `status` BIGINT NOT NULL DEFAULT 0,
    `sort_order` FLOAT NOT NULL DEFAULT 0,
    `queued_at` BIGINT NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `payload` JSON NULL,
    PRIMARY KEY (`id`),
    INDEX (`status`),
    CONSTRAINT `fk_categories_to_category_types` FOREIGN KEY (`category_type_id`) REFERENCES `category_types` (`id`),
    CONSTRAINT `fk_categories_to_categories` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  COLLATE = `utf8mb4_unicode_ci`;

DELIMITER $$

CREATE TRIGGER `categories_after_update`
AFTER UPDATE
ON `categories` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`categories`(`action`, `original_id`, `category_type_id`, `category_id`, `title`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`category_type_id`, `old`.`category_id`, `old`.`title`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `categories_after_delete`
AFTER DELETE
ON `categories` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`categories`(`action`, `original_id`, `category_type_id`, `category_id`, `title`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`category_type_id`, `old`.`category_id`, `old`.`title`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Users
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `users`
(
    `id` BIGINT NOT NULL,
    `github` VARCHAR(512) NOT NULL,
    `editor` BIGINT NOT NULL DEFAULT 0,
    `status` BIGINT NOT NULL DEFAULT 0,
    `sort_order` FLOAT NOT NULL DEFAULT 0,
    `queued_at` BIGINT NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `payload` JSON NULL,
    PRIMARY KEY (`id`),
    INDEX (`status`),
    CONSTRAINT `fk_users_to_identities` FOREIGN KEY (`id`) REFERENCES `identities` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  COLLATE = `utf8mb4_unicode_ci`;

DELIMITER $$

CREATE TRIGGER `users_after_update`
AFTER UPDATE
ON `users` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`users`(`action`, `original_id`, `github`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`github`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `users_after_delete`
AFTER DELETE
ON `users` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`users`(`action`, `original_id`, `github`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`github`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	ActivityPubIncomingActivities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `activity_pub_incoming_activities`
(
    `id` BIGINT NOT NULL,
    `identity_id` BIGINT NOT NULL,
    `unique_identifier` VARCHAR(128) NOT NULL,
    `timestamp` BIGINT NOT NULL,
    `from` VARCHAR(256) NOT NULL,
    `to` VARCHAR(256) NOT NULL,
    `content` VARCHAR(4096) NOT NULL,
    `raw` JSON NOT NULL,
    `editor` BIGINT NOT NULL DEFAULT 0,
    `status` BIGINT NOT NULL DEFAULT 0,
    `sort_order` FLOAT NOT NULL DEFAULT 0,
    `queued_at` BIGINT NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `payload` JSON NULL,
    PRIMARY KEY (`id`),
    INDEX (`status`),
    CONSTRAINT `fk_activity_pub_incoming_activities_to_identities` FOREIGN KEY (`identity_id`) REFERENCES `identities` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  COLLATE = `utf8mb4_unicode_ci`;

DELIMITER $$

CREATE TRIGGER `activity_pub_incoming_activities_after_update`
AFTER UPDATE
ON `activity_pub_incoming_activities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`activity_pub_incoming_activities`(`action`, `original_id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`identity_id`, `old`.`unique_identifier`, `old`.`timestamp`, `old`.`from`, `old`.`to`, `old`.`content`, `old`.`raw`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `activity_pub_incoming_activities_after_delete`
AFTER DELETE
ON `activity_pub_incoming_activities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`activity_pub_incoming_activities`(`action`, `original_id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`identity_id`, `old`.`unique_identifier`, `old`.`timestamp`, `old`.`from`, `old`.`to`, `old`.`content`, `old`.`raw`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	ActivityPubOutgoingActivities
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

CREATE TABLE `activity_pub_outgoing_activities`
(
    `id` BIGINT NOT NULL,
    `identity_id` BIGINT NOT NULL,
    `unique_identifier` VARCHAR(128) NOT NULL,
    `timestamp` BIGINT NOT NULL,
    `from` VARCHAR(256) NOT NULL,
    `to` VARCHAR(256) NOT NULL,
    `content` VARCHAR(4096) NOT NULL,
    `raw` JSON NOT NULL,
    `editor` BIGINT NOT NULL DEFAULT 0,
    `status` BIGINT NOT NULL DEFAULT 0,
    `sort_order` FLOAT NOT NULL DEFAULT 0,
    `queued_at` BIGINT NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `payload` JSON NULL,
    PRIMARY KEY (`id`),
    INDEX (`status`),
    CONSTRAINT `fk_activity_pub_outgoing_activities_to_identities` FOREIGN KEY (`identity_id`) REFERENCES `identities` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  COLLATE = `utf8mb4_unicode_ci`;

DELIMITER $$

CREATE TRIGGER `activity_pub_outgoing_activities_after_update`
AFTER UPDATE
ON `activity_pub_outgoing_activities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`activity_pub_outgoing_activities`(`action`, `original_id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`identity_id`, `old`.`unique_identifier`, `old`.`timestamp`, `old`.`from`, `old`.`to`, `old`.`content`, `old`.`raw`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

CREATE TRIGGER `activity_pub_outgoing_activities_after_delete`
AFTER DELETE
ON `activity_pub_outgoing_activities` FOR EACH ROW
BEGIN
    INSERT INTO `greatape_dev_history`.`activity_pub_outgoing_activities`(`action`, `original_id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`identity_id`, `old`.`unique_identifier`, `old`.`timestamp`, `old`.`from`, `old`.`to`, `old`.`content`, `old`.`raw`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END$$

DELIMITER ;

# ╔═════════════════════════════════════════════════════════════════════════════════════════════════════════════
# ║	Initialization
# ╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════

# Identities
INSERT INTO `identities` (`id`, `username`, `phone_number`, `phone_number_confirmed`, `first_name`, `last_name`, `display_name`, `email`, `email_confirmed`, `avatar`, `banner`, `summary`, `token`, `multi_factor`, `hash`, `salt`, `public_key`, `private_key`, `permission`, `restriction`, `last_login`, `login_count`) VALUES (0, 'INVALID', '0', false, '', '', '', 'invalid@localhost', false, '', '', '', '0', b'0', '', '', '', '', 0, 0, 0, 0);
INSERT INTO `identities` (`id`, `username`, `phone_number`, `phone_number_confirmed`, `first_name`, `last_name`, `display_name`, `email`, `email_confirmed`, `avatar`, `banner`, `summary`, `token`, `multi_factor`, `hash`, `salt`, `public_key`, `private_key`, `permission`, `restriction`, `last_login`, `login_count`) VALUES (1, 'root', '1', false, '', '', '', 'root@localhost', false, '', '', '', '1', b'1', '', '', '', '', 0xFFFFFFFF, 0, 0, 0);

# Users
INSERT INTO `users` (`id`, `github`) VALUES (0, '');
INSERT INTO `users` (`id`, `github`) VALUES (1, '');

# CategoryTypes
INSERT INTO `category_types` (`id`, `description`) VALUES (0, 'INVALID');

# Categories
INSERT INTO `categories` (`id`, `category_type_id`, `category_id`, `title`, `description`) VALUES (0, 0, 0, 'INVALID', '');
