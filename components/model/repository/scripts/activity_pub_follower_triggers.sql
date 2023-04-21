
USE `###DATABASE###`;

CREATE TRIGGER `activity_pub_followers_after_update`
AFTER UPDATE
ON `activity_pub_followers` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`activity_pub_followers`(`action`, `original_id`, `handle`, `inbox`, `subject`, `activity`, `accepted`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`handle`, `old`.`inbox`, `old`.`subject`, `old`.`activity`, `old`.`accepted`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `activity_pub_followers_after_delete`
AFTER DELETE
ON `activity_pub_followers` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`activity_pub_followers`(`action`, `original_id`, `handle`, `inbox`, `subject`, `activity`, `accepted`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`handle`, `old`.`inbox`, `old`.`subject`, `old`.`activity`, `old`.`accepted`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
