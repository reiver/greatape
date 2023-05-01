########## 

CREATE TABLE "system_schedules_history"
(
    "id" BIGSERIAL NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "original_id" BIGINT NOT NULL,
    "enabled" BOOLEAN NOT NULL,
    "config" VARCHAR(1024) NOT NULL,
    "editor" BIGINT NOT NULL,
    "status" BIGINT NOT NULL,
    "sort_order" REAL NOT NULL,
    "queued_at" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "triggered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "changed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_system_schedules_history" PRIMARY KEY ("id")
);

##########

CREATE TABLE "system_schedules"
(
    "id" BIGINT NOT NULL,
    "enabled" BOOLEAN NOT NULL,
    "config" VARCHAR(1024) NOT NULL,
    "editor" BIGINT NOT NULL DEFAULT 0,
    "status" BIGINT NOT NULL DEFAULT 0,
    "sort_order" REAL NOT NULL DEFAULT 0,
    "queued_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_system_schedules" PRIMARY KEY ("id")
);

##########

CREATE INDEX "idx_system_schedules_status" ON "system_schedules" ("status");

