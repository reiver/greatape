
USE `###DATABASE###`;

CREATE TRIGGER `category_types_after_update`
AFTER UPDATE
ON `category_types` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`category_types`(`action`, `original_id`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `category_types_after_delete`
AFTER DELETE
ON `category_types` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`category_types`(`action`, `original_id`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
