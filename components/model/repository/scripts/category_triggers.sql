########## 

CREATE OR REPLACE FUNCTION "categories_before_update"() RETURNS TRIGGER AS $categories_before_update$
    BEGIN
        INSERT INTO "categories_history"("action", "original_id", "category_type_id", "category_id", "title", "description", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('update', OLD."id", OLD."category_type_id", OLD."category_id", OLD."title", OLD."description", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN NEW;
    END;
$categories_before_update$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "categories_before_update_trigger" BEFORE UPDATE ON "categories"
    FOR EACH ROW EXECUTE FUNCTION "categories_before_update"();

##########

CREATE OR REPLACE FUNCTION "categories_before_delete"() RETURNS TRIGGER AS $categories_before_delete$
    BEGIN
        INSERT INTO "categories_history"("action", "original_id", "category_type_id", "category_id", "title", "description", "editor", "status", "sort_order", "queued_at", "created_at", "updated_at", "payload")
        VALUES('delete', OLD."id", OLD."category_type_id", OLD."category_id", OLD."title", OLD."description", OLD."editor", OLD."status", OLD."sort_order", OLD."queued_at", OLD."created_at", OLD."updated_at", OLD."payload");
        RETURN OLD;
    END;
$categories_before_delete$ LANGUAGE plpgsql;

##########

CREATE OR REPLACE TRIGGER "categories_before_delete_trigger" BEFORE DELETE ON "categories"
    FOR EACH ROW EXECUTE FUNCTION "categories_before_delete"();

##########
