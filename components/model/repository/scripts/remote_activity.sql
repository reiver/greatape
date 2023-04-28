########## 

CREATE TABLE "remote_activities_history"
(
    "id" BIGSERIAL NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "original_id" BIGINT NOT NULL,
    "entry_point" VARCHAR(256) NOT NULL,
    "duration" BIGINT NOT NULL,
    "successful" BOOLEAN NOT NULL,
    "error_message" VARCHAR(1024) NOT NULL,
    "remote_address" VARCHAR(128) NOT NULL,
    "user_agent" VARCHAR(512) NOT NULL,
    "event_type" INT NOT NULL,
    "timestamp" BIGINT NOT NULL,
    "editor" BIGINT NOT NULL,
    "status" BIGINT NOT NULL,
    "sort_order" REAL NOT NULL,
    "queued_at" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "triggered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "changed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_remote_activities_history" PRIMARY KEY ("id")
);

##########

CREATE TABLE "remote_activities"
(
    "id" BIGINT NOT NULL,
    "entry_point" VARCHAR(256) NOT NULL,
    "duration" BIGINT NOT NULL,
    "successful" BOOLEAN NOT NULL,
    "error_message" VARCHAR(1024) NOT NULL,
    "remote_address" VARCHAR(128) NOT NULL,
    "user_agent" VARCHAR(512) NOT NULL,
    "event_type" INT NOT NULL,
    "timestamp" BIGINT NOT NULL,
    "editor" BIGINT NOT NULL DEFAULT 0,
    "status" BIGINT NOT NULL DEFAULT 0,
    "sort_order" REAL NOT NULL DEFAULT 0,
    "queued_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_remote_activities" PRIMARY KEY ("id")
);

##########

CREATE INDEX "idx_remote_activities_status" ON "remote_activities" ("status");
