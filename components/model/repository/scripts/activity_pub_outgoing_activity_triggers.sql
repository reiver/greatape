########## 

CREATE OR REPLACE FUNCTION "activity_pub_outgoing_activities_before_update"() RETURNS TRIGGER AS $activity_pub_outgoing_activities_before_update$
    BEGIN
        INSERT INTO "activity_pub_outgoing_activities_history"("action", "original_id", "identity_id", "unique_identifier", "timestamp", "from", "to", "content", "raw", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', OLD."id", OLD."identity_id", OLD."unique_identifier", OLD."timestamp", OLD."from", OLD."to", OLD."content", OLD."raw", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN NEW;
    END;
$activity_pub_outgoing_activities_before_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "activity_pub_outgoing_activities_before_update_trigger" BEFORE UPDATE ON "activity_pub_outgoing_activities"
    FOR EACH ROW EXECUTE FUNCTION "activity_pub_outgoing_activities_before_update"();

##########

CREATE OR REPLACE FUNCTION "activity_pub_outgoing_activities_before_delete"() RETURNS TRIGGER AS $activity_pub_outgoing_activities_before_delete$
    BEGIN
        INSERT INTO "activity_pub_outgoing_activities_history"("action", "original_id", "identity_id", "unique_identifier", "timestamp", "from", "to", "content", "raw", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', OLD."id", OLD."identity_id", OLD."unique_identifier", OLD."timestamp", OLD."from", OLD."to", OLD."content", OLD."raw", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN OLD;
    END;
$activity_pub_outgoing_activities_before_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "activity_pub_outgoing_activities_before_delete_trigger" BEFORE DELETE ON "activity_pub_outgoing_activities"
    FOR EACH ROW EXECUTE FUNCTION "activity_pub_outgoing_activities_before_delete"();

##########
