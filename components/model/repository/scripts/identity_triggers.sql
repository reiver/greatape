########## 

CREATE OR REPLACE FUNCTION "identities_after_update"() RETURNS TRIGGER AS $identities_after_update$
    BEGIN
        INSERT INTO "identities_history"("action", "original_id", "username", "phone_number", "phone_number_confirmed", "first_name", "last_name", "display_name", "email", "email_confirmed", "avatar", "banner", "summary", "token", "multi_factor", "hash", "salt", "public_key", "private_key", "permission", "restriction", "last_login", "login_count", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', "OLD"."id", "OLD"."username", "OLD"."phone_number", "OLD"."phone_number_confirmed", "OLD"."first_name", "OLD"."last_name", "OLD"."display_name", "OLD"."email", "OLD"."email_confirmed", "OLD"."avatar", "OLD"."banner", "OLD"."summary", "OLD"."token", "OLD"."multi_factor", "OLD"."hash", "OLD"."salt", "OLD"."public_key", "OLD"."private_key", "OLD"."permission", "OLD"."restriction", "OLD"."last_login", "OLD"."login_count", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$identities_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "identities_after_update_trigger" AFTER UPDATE ON "identities"
    FOR EACH ROW EXECUTE FUNCTION "identities_after_update"();

##########

CREATE OR REPLACE FUNCTION "identities_after_delete"() RETURNS TRIGGER AS $identities_after_delete$
    BEGIN
        INSERT INTO "identities_history"("action", "original_id", "username", "phone_number", "phone_number_confirmed", "first_name", "last_name", "display_name", "email", "email_confirmed", "avatar", "banner", "summary", "token", "multi_factor", "hash", "salt", "public_key", "private_key", "permission", "restriction", "last_login", "login_count", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', "OLD"."id", "OLD"."username", "OLD"."phone_number", "OLD"."phone_number_confirmed", "OLD"."first_name", "OLD"."last_name", "OLD"."display_name", "OLD"."email", "OLD"."email_confirmed", "OLD"."avatar", "OLD"."banner", "OLD"."summary", "OLD"."token", "OLD"."multi_factor", "OLD"."hash", "OLD"."salt", "OLD"."public_key", "OLD"."private_key", "OLD"."permission", "OLD"."restriction", "OLD"."last_login", "OLD"."login_count", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$identities_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "identities_after_delete_trigger" AFTER DELETE ON "identities"
    FOR EACH ROW EXECUTE FUNCTION "identities_after_delete"();

