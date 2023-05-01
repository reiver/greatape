########## 

CREATE TABLE "activity_pub_incoming_activities_history"
(
    "id" BIGSERIAL NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "original_id" BIGINT NOT NULL,
    "identity_id" BIGINT NOT NULL,
    "unique_identifier" VARCHAR(128) NOT NULL,
    "timestamp" BIGINT NOT NULL,
    "from" VARCHAR(256) NOT NULL,
    "to" VARCHAR(256) NOT NULL,
    "content" VARCHAR(4096) NOT NULL,
    "raw" JSONB NOT NULL,
    "editor" BIGINT NOT NULL,
    "status" BIGINT NOT NULL,
    "sort_order" REAL NOT NULL,
    "queued_at" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "triggered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "changed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_activity_pub_incoming_activities_history" PRIMARY KEY ("id")
);

##########

CREATE TABLE "activity_pub_incoming_activities"
(
    "id" BIGINT NOT NULL,
    "identity_id" BIGINT NOT NULL,
    "unique_identifier" VARCHAR(128) NOT NULL,
    "timestamp" BIGINT NOT NULL,
    "from" VARCHAR(256) NOT NULL,
    "to" VARCHAR(256) NOT NULL,
    "content" VARCHAR(4096) NOT NULL,
    "raw" JSONB NOT NULL,
    "editor" BIGINT NOT NULL DEFAULT 0,
    "status" BIGINT NOT NULL DEFAULT 0,
    "sort_order" REAL NOT NULL DEFAULT 0,
    "queued_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_activity_pub_incoming_activities" PRIMARY KEY ("id"),
    CONSTRAINT "fk_activity_pub_incoming_activities_to_identities" FOREIGN KEY ("identity_id") REFERENCES "identities" ("id")
);

##########

CREATE INDEX "idx_activity_pub_incoming_activities_status" ON "activity_pub_incoming_activities" ("status");

