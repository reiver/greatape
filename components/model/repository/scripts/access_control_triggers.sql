
USE `###DATABASE###`;

CREATE TRIGGER `access_controls_after_update`
AFTER UPDATE
ON `access_controls` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`access_controls`(`action`, `original_id`, `key`, `value`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`key`, `old`.`value`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `access_controls_after_delete`
AFTER DELETE
ON `access_controls` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`access_controls`(`action`, `original_id`, `key`, `value`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`key`, `old`.`value`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
