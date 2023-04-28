########## 

CREATE OR REPLACE FUNCTION "category_types_after_update"() RETURNS TRIGGER AS $category_types_after_update$
    BEGIN
        INSERT INTO "category_types_history"("action", "original_id", "description", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', "OLD"."id", "OLD"."description", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$category_types_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "category_types_after_update_trigger" AFTER UPDATE ON "category_types"
    FOR EACH ROW EXECUTE FUNCTION "category_types_after_update"();

##########

CREATE OR REPLACE FUNCTION "category_types_after_delete"() RETURNS TRIGGER AS $category_types_after_delete$
    BEGIN
        INSERT INTO "category_types_history"("action", "original_id", "description", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', "OLD"."id", "OLD"."description", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$category_types_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "category_types_after_delete_trigger" AFTER DELETE ON "category_types"
    FOR EACH ROW EXECUTE FUNCTION "category_types_after_delete"();

