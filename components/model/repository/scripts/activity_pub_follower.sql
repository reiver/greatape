########## 

CREATE TABLE "activity_pub_followers_history"
(
    "id" BIGSERIAL NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "original_id" BIGINT NOT NULL,
    "handle" VARCHAR(256) NOT NULL,
    "inbox" VARCHAR(256) NOT NULL,
    "subject" VARCHAR(256) NOT NULL,
    "activity" VARCHAR(4096) NOT NULL,
    "accepted" BOOLEAN NOT NULL,
    "editor" BIGINT NOT NULL,
    "status" BIGINT NOT NULL,
    "sort_order" REAL NOT NULL,
    "queued_at" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "triggered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "changed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_activity_pub_followers_history" PRIMARY KEY ("id")
);

##########

CREATE TABLE "activity_pub_followers"
(
    "id" BIGINT NOT NULL,
    "handle" VARCHAR(256) NOT NULL,
    "inbox" VARCHAR(256) NOT NULL,
    "subject" VARCHAR(256) NOT NULL,
    "activity" VARCHAR(4096) NOT NULL,
    "accepted" BOOLEAN NOT NULL,
    "editor" BIGINT NOT NULL DEFAULT 0,
    "status" BIGINT NOT NULL DEFAULT 0,
    "sort_order" REAL NOT NULL DEFAULT 0,
    "queued_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_activity_pub_followers" PRIMARY KEY ("id")
);

##########

CREATE INDEX "idx_activity_pub_followers_status" ON "activity_pub_followers" ("status");
