########## 

CREATE TABLE "access_controls_history"
(
    "id" BIGSERIAL NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "original_id" BIGINT NOT NULL,
    "key" BIGINT NOT NULL,
    "value" BIGINT NOT NULL,
    "editor" BIGINT NOT NULL,
    "status" BIGINT NOT NULL,
    "sort_order" REAL NOT NULL,
    "queued_at" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "triggered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "changed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_access_controls_history" PRIMARY KEY ("id")
);

##########

CREATE TABLE "access_controls"
(
    "id" BIGINT NOT NULL,
    "key" BIGINT NOT NULL,
    "value" BIGINT NOT NULL,
    "editor" BIGINT NOT NULL DEFAULT 0,
    "status" BIGINT NOT NULL DEFAULT 0,
    "sort_order" REAL NOT NULL DEFAULT 0,
    "queued_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_access_controls" PRIMARY KEY ("id")
);

##########

CREATE INDEX "idx_access_controls_status" ON "access_controls" ("status");

