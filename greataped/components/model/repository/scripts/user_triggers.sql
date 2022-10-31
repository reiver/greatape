
USE `###DATABASE###`;

CREATE TRIGGER `users_after_update`
AFTER UPDATE
ON `users` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`users`(`action`, `original_id`, `github`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`github`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `users_after_delete`
AFTER DELETE
ON `users` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`users`(`action`, `original_id`, `github`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`github`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
