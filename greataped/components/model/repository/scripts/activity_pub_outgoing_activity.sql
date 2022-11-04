USE `###DATABASE###_history`;

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

USE `###DATABASE###`;

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
