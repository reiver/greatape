
USE `###DATABASE###`;

CREATE TRIGGER `identities_after_update`
AFTER UPDATE
ON `identities` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`identities`(`action`, `original_id`, `username`, `phone_number`, `phone_number_confirmed`, `first_name`, `last_name`, `display_name`, `email`, `email_confirmed`, `avatar`, `banner`, `summary`, `token`, `multi_factor`, `hash`, `salt`, `public_key`, `private_key`, `permission`, `restriction`, `last_login`, `login_count`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('update', `old`.`id`, `old`.`username`, `old`.`phone_number`, `old`.`phone_number_confirmed`, `old`.`first_name`, `old`.`last_name`, `old`.`display_name`, `old`.`email`, `old`.`email_confirmed`, `old`.`avatar`, `old`.`banner`, `old`.`summary`, `old`.`token`, `old`.`multi_factor`, `old`.`hash`, `old`.`salt`, `old`.`public_key`, `old`.`private_key`, `old`.`permission`, `old`.`restriction`, `old`.`last_login`, `old`.`login_count`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;

CREATE TRIGGER `identities_after_delete`
AFTER DELETE
ON `identities` FOR EACH ROW
BEGIN
    INSERT INTO `###DATABASE###_history`.`identities`(`action`, `original_id`, `username`, `phone_number`, `phone_number_confirmed`, `first_name`, `last_name`, `display_name`, `email`, `email_confirmed`, `avatar`, `banner`, `summary`, `token`, `multi_factor`, `hash`, `salt`, `public_key`, `private_key`, `permission`, `restriction`, `last_login`, `login_count`, `editor`, `status`, `sort_order`, `queued_at`, `created_at`, `updated_at`, `payload`)
    VALUES('delete', `old`.`id`, `old`.`username`, `old`.`phone_number`, `old`.`phone_number_confirmed`, `old`.`first_name`, `old`.`last_name`, `old`.`display_name`, `old`.`email`, `old`.`email_confirmed`, `old`.`avatar`, `old`.`banner`, `old`.`summary`, `old`.`token`, `old`.`multi_factor`, `old`.`hash`, `old`.`salt`, `old`.`public_key`, `old`.`private_key`, `old`.`permission`, `old`.`restriction`, `old`.`last_login`, `old`.`login_count`, `old`.`editor`, `old`.`status`, `old`.`sort_order`, `old`.`queued_at`, `old`.`created_at`, `old`.`updated_at`, `old`.`payload`);
END;
