########## 

CREATE OR REPLACE FUNCTION "system_schedules_before_update"() RETURNS TRIGGER AS $system_schedules_before_update$
    BEGIN
        INSERT INTO "system_schedules_history"("action", "original_id", "enabled", "config", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', OLD."id", OLD."enabled", OLD."config", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN NEW;
    END;
$system_schedules_before_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "system_schedules_before_update_trigger" BEFORE UPDATE ON "system_schedules"
    FOR EACH ROW EXECUTE FUNCTION "system_schedules_before_update"();

##########

CREATE OR REPLACE FUNCTION "system_schedules_before_delete"() RETURNS TRIGGER AS $system_schedules_before_delete$
    BEGIN
        INSERT INTO "system_schedules_history"("action", "original_id", "enabled", "config", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', OLD."id", OLD."enabled", OLD."config", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN OLD;
    END;
$system_schedules_before_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "system_schedules_before_delete_trigger" BEFORE DELETE ON "system_schedules"
    FOR EACH ROW EXECUTE FUNCTION "system_schedules_before_delete"();

##########
