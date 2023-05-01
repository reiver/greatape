########## 

CREATE OR REPLACE FUNCTION "activity_pub_followers_after_update"() RETURNS TRIGGER AS $activity_pub_followers_after_update$
    BEGIN
        INSERT INTO "activity_pub_followers_history"("action", "original_id", "handle", "inbox", "subject", "activity", "accepted", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', OLD."id", OLD."handle", OLD."inbox", OLD."subject", OLD."activity", OLD."accepted", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN NEW;
    END;
$activity_pub_followers_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "activity_pub_followers_after_update_trigger" BEFORE UPDATE ON "activity_pub_followers"
    FOR EACH ROW EXECUTE FUNCTION "activity_pub_followers_after_update"();

##########

CREATE OR REPLACE FUNCTION "activity_pub_followers_after_delete"() RETURNS TRIGGER AS $activity_pub_followers_after_delete$
    BEGIN
        INSERT INTO "activity_pub_followers_history"("action", "original_id", "handle", "inbox", "subject", "activity", "accepted", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', OLD."id", OLD."handle", OLD."inbox", OLD."subject", OLD."activity", OLD."accepted", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN OLD;
    END;
$activity_pub_followers_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "activity_pub_followers_after_delete_trigger" BEFORE DELETE ON "activity_pub_followers"
    FOR EACH ROW EXECUTE FUNCTION "activity_pub_followers_after_delete"();

##########
