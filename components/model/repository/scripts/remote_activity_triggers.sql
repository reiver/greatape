########## 

CREATE OR REPLACE FUNCTION "remote_activities_after_update"() RETURNS TRIGGER AS $remote_activities_after_update$
    BEGIN
        INSERT INTO "remote_activities_history"("action", "original_id", "entry_point", "duration", "successful", "error_message", "remote_address", "user_agent", "event_type", "timestamp", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', OLD."id", OLD."entry_point", OLD."duration", OLD."successful", OLD."error_message", OLD."remote_address", OLD."user_agent", OLD."event_type", OLD."timestamp", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN NEW;
    END;
$remote_activities_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "remote_activities_after_update_trigger" BEFORE UPDATE ON "remote_activities"
    FOR EACH ROW EXECUTE FUNCTION "remote_activities_after_update"();

##########

CREATE OR REPLACE FUNCTION "remote_activities_after_delete"() RETURNS TRIGGER AS $remote_activities_after_delete$
    BEGIN
        INSERT INTO "remote_activities_history"("action", "original_id", "entry_point", "duration", "successful", "error_message", "remote_address", "user_agent", "event_type", "timestamp", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', OLD."id", OLD."entry_point", OLD."duration", OLD."successful", OLD."error_message", OLD."remote_address", OLD."user_agent", OLD."event_type", OLD."timestamp", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN OLD;
    END;
$remote_activities_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "remote_activities_after_delete_trigger" BEFORE DELETE ON "remote_activities"
    FOR EACH ROW EXECUTE FUNCTION "remote_activities_after_delete"();

##########
