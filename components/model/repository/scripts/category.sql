########## 

CREATE TABLE "categories_history"
(
    "id" BIGSERIAL NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "original_id" BIGINT NOT NULL,
    "category_type_id" BIGINT NOT NULL,
    "category_id" BIGINT NOT NULL,
    "title" VARCHAR(64) NOT NULL,
    "description" VARCHAR(64) NOT NULL,
    "editor" BIGINT NOT NULL,
    "status" BIGINT NOT NULL,
    "sort_order" REAL NOT NULL,
    "queued_at" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "triggered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "changed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_categories_history" PRIMARY KEY ("id")
);

##########

CREATE TABLE "categories"
(
    "id" BIGINT NOT NULL,
    "category_type_id" BIGINT NOT NULL,
    "category_id" BIGINT NOT NULL,
    "title" VARCHAR(64) NOT NULL,
    "description" VARCHAR(64) NOT NULL,
    "editor" BIGINT NOT NULL DEFAULT 0,
    "status" BIGINT NOT NULL DEFAULT 0,
    "sort_order" REAL NOT NULL DEFAULT 0,
    "queued_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_categories" PRIMARY KEY ("id"),
    CONSTRAINT "fk_categories_to_category_types" FOREIGN KEY ("category_type_id") REFERENCES "category_types" ("id"),
    CONSTRAINT "fk_categories_to_categories" FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
);

##########

CREATE INDEX "idx_categories_status" ON "categories" ("status");
