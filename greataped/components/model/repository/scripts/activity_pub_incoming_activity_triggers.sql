
USE `###DATABASE###`;

CREATE TRIGGER `activity_pub_incoming_activities_after_update`
AFTER UPDATE
ON `activity_pub_incoming_activities` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`activity_pub_incoming_activities`(`action`, `original_id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`identity_id`, `old`.`unique_identifier`, `old`.`timestamp`, `old`.`from`, `old`.`to`, `old`.`content`, `old`.`raw`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `activity_pub_incoming_activities_after_delete`
AFTER DELETE
ON `activity_pub_incoming_activities` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`activity_pub_incoming_activities`(`action`, `original_id`, `identity_id`, `unique_identifier`, `timestamp`, `from`, `to`, `content`, `raw`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`identity_id`, `old`.`unique_identifier`, `old`.`timestamp`, `old`.`from`, `old`.`to`, `old`.`content`, `old`.`raw`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
