
USE `###DATABASE###`;

CREATE TRIGGER `categories_after_update`
AFTER UPDATE
ON `categories` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`categories`(`action`, `original_id`, `category_type_id`, `category_id`, `title`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`category_type_id`, `old`.`category_id`, `old`.`title`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `categories_after_delete`
AFTER DELETE
ON `categories` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`categories`(`action`, `original_id`, `category_type_id`, `category_id`, `title`, `description`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`category_type_id`, `old`.`category_id`, `old`.`title`, `old`.`description`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
