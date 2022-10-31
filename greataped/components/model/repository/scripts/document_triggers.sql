
USE `###DATABASE###`;

CREATE TRIGGER `documents_after_update`
AFTER UPDATE
ON `documents` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`documents`(`action`, `original_id`, `content`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`content`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `documents_after_delete`
AFTER DELETE
ON `documents` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`documents`(`action`, `original_id`, `content`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`content`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
