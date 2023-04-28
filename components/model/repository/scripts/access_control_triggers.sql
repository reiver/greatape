########## 

CREATE OR REPLACE FUNCTION "access_controls_after_update"() RETURNS TRIGGER AS $access_controls_after_update$
    BEGIN
        INSERT INTO "access_controls_history"("action", "original_id", "key", "value", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', "OLD"."id", "OLD"."key", "OLD"."value", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$access_controls_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "access_controls_after_update_trigger" AFTER UPDATE ON "access_controls"
    FOR EACH ROW EXECUTE FUNCTION "access_controls_after_update"();

##########

CREATE OR REPLACE FUNCTION "access_controls_after_delete"() RETURNS TRIGGER AS $access_controls_after_delete$
    BEGIN
        INSERT INTO "access_controls_history"("action", "original_id", "key", "value", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', "OLD"."id", "OLD"."key", "OLD"."value", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$access_controls_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "access_controls_after_delete_trigger" AFTER DELETE ON "access_controls"
    FOR EACH ROW EXECUTE FUNCTION "access_controls_after_delete"();

