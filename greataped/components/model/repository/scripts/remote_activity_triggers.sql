
USE `###DATABASE###`;

CREATE TRIGGER `remote_activities_after_update`
AFTER UPDATE
ON `remote_activities` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`remote_activities`(`action`, `original_id`, `entry_point`, `duration`, `successful`, `error_message`, `remote_address`, `user_agent`, `event_type`, `timestamp`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`entry_point`, `old`.`duration`, `old`.`successful`, `old`.`error_message`, `old`.`remote_address`, `old`.`user_agent`, `old`.`event_type`, `old`.`timestamp`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `remote_activities_after_delete`
AFTER DELETE
ON `remote_activities` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`remote_activities`(`action`, `original_id`, `entry_point`, `duration`, `successful`, `error_message`, `remote_address`, `user_agent`, `event_type`, `timestamp`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`entry_point`, `old`.`duration`, `old`.`successful`, `old`.`error_message`, `old`.`remote_address`, `old`.`user_agent`, `old`.`event_type`, `old`.`timestamp`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
