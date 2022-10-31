
USE `###DATABASE###`;

CREATE TRIGGER `system_schedules_after_update`
AFTER UPDATE
ON `system_schedules` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`system_schedules`(`action`, `original_id`, `enabled`, `config`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`enabled`, `old`.`config`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `system_schedules_after_delete`
AFTER DELETE
ON `system_schedules` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`system_schedules`(`action`, `original_id`, `enabled`, `config`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`enabled`, `old`.`config`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
