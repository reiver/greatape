########## 

CREATE OR REPLACE FUNCTION "documents_after_update"() RETURNS TRIGGER AS $documents_after_update$
    BEGIN
        INSERT INTO "documents_history"("action", "original_id", "content", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', OLD."id", OLD."content", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN NEW;
    END;
$documents_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "documents_after_update_trigger" BEFORE UPDATE ON "documents"
    FOR EACH ROW EXECUTE FUNCTION "documents_after_update"();

##########

CREATE OR REPLACE FUNCTION "documents_after_delete"() RETURNS TRIGGER AS $documents_after_delete$
    BEGIN
        INSERT INTO "documents_history"("action", "original_id", "content", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', OLD."id", OLD."content", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN OLD;
    END;
$documents_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "documents_after_delete_trigger" BEFORE DELETE ON "documents"
    FOR EACH ROW EXECUTE FUNCTION "documents_after_delete"();

##########
