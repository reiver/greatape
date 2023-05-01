########## 

CREATE OR REPLACE FUNCTION "users_after_update"() RETURNS TRIGGER AS $users_after_update$
    BEGIN
        INSERT INTO "users_history"("action", "original_id", "github", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', OLD."id", OLD."github", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN NEW;
    END;
$users_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "users_after_update_trigger" BEFORE UPDATE ON "users"
    FOR EACH ROW EXECUTE FUNCTION "users_after_update"();

##########

CREATE OR REPLACE FUNCTION "users_after_delete"() RETURNS TRIGGER AS $users_after_delete$
    BEGIN
        INSERT INTO "users_history"("action", "original_id", "github", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', OLD."id", OLD."github", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN OLD;
    END;
$users_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "users_after_delete_trigger" BEFORE DELETE ON "users"
    FOR EACH ROW EXECUTE FUNCTION "users_after_delete"();

##########
