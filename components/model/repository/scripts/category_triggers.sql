########## 

CREATE OR REPLACE FUNCTION "categories_after_update"() RETURNS TRIGGER AS $categories_after_update$
    BEGIN
        INSERT INTO "categories_history"("action", "original_id", "category_type_id", "category_id", "title", "description", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', "OLD"."id", "OLD"."category_type_id", "OLD"."category_id", "OLD"."title", "OLD"."description", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$categories_after_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "categories_after_update_trigger" AFTER UPDATE ON "categories"
    FOR EACH ROW EXECUTE FUNCTION "categories_after_update"();

##########

CREATE OR REPLACE FUNCTION "categories_after_delete"() RETURNS TRIGGER AS $categories_after_delete$
    BEGIN
        INSERT INTO "categories_history"("action", "original_id", "category_type_id", "category_id", "title", "description", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', "OLD"."id", "OLD"."category_type_id", "OLD"."category_id", "OLD"."title", "OLD"."description", "OLD"."editor", "OLD"."status", "OLD"."sort_order", "OLD"."queued_at", "OLD"."created_at", "OLD"."updated_at", "OLD"."payload");
    END;
$categories_after_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "categories_after_delete_trigger" AFTER DELETE ON "categories"
    FOR EACH ROW EXECUTE FUNCTION "categories_after_delete"();

