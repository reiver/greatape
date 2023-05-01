########## 

CREATE TABLE "identities_history"
(
    "id" BIGSERIAL NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "original_id" BIGINT NOT NULL,
    "username" VARCHAR(32) NOT NULL,
    "phone_number" VARCHAR(12) NOT NULL,
    "phone_number_confirmed" BOOLEAN NOT NULL,
    "first_name" VARCHAR(128) NOT NULL,
    "last_name" VARCHAR(128) NOT NULL,
    "display_name" VARCHAR(128) NOT NULL,
    "email" VARCHAR(128) NOT NULL,
    "email_confirmed" BOOLEAN NOT NULL,
    "avatar" VARCHAR(512) NOT NULL,
    "banner" VARCHAR(512) NOT NULL,
    "summary" VARCHAR(512) NOT NULL,
    "token" VARCHAR(256) NOT NULL,
    "multi_factor" BOOLEAN NOT NULL,
    "hash" VARCHAR(256) NOT NULL,
    "salt" VARCHAR(64) NOT NULL,
    "public_key" VARCHAR(4096) NOT NULL,
    "private_key" VARCHAR(4096) NOT NULL,
    "permission" BIGINT NOT NULL,
    "restriction" INT NOT NULL,
    "last_login" BIGINT NOT NULL,
    "login_count" INT NOT NULL,
    "editor" BIGINT NOT NULL,
    "status" BIGINT NOT NULL,
    "sort_order" REAL NOT NULL,
    "queued_at" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "triggered_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "changed_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_identities_history" PRIMARY KEY ("id")
);

##########

CREATE TABLE "identities"
(
    "id" BIGINT NOT NULL,
    "username" VARCHAR(32) NOT NULL CONSTRAINT "udx_identities_username" UNIQUE,
    "phone_number" VARCHAR(12) NOT NULL,
    "phone_number_confirmed" BOOLEAN NOT NULL,
    "first_name" VARCHAR(128) NOT NULL,
    "last_name" VARCHAR(128) NOT NULL,
    "display_name" VARCHAR(128) NOT NULL,
    "email" VARCHAR(128) NOT NULL CONSTRAINT "udx_identities_email" UNIQUE,
    "email_confirmed" BOOLEAN NOT NULL,
    "avatar" VARCHAR(512) NOT NULL,
    "banner" VARCHAR(512) NOT NULL,
    "summary" VARCHAR(512) NOT NULL,
    "token" VARCHAR(256) NOT NULL CONSTRAINT "udx_identities_token" UNIQUE,
    "multi_factor" BOOLEAN NOT NULL,
    "hash" VARCHAR(256) NOT NULL,
    "salt" VARCHAR(64) NOT NULL,
    "public_key" VARCHAR(4096) NOT NULL,
    "private_key" VARCHAR(4096) NOT NULL,
    "permission" BIGINT NOT NULL,
    "restriction" INT NOT NULL,
    "last_login" BIGINT NOT NULL,
    "login_count" INT NOT NULL,
    "editor" BIGINT NOT NULL DEFAULT 0,
    "status" BIGINT NOT NULL DEFAULT 0,
    "sort_order" REAL NOT NULL DEFAULT 0,
    "queued_at" BIGINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "payload" JSONB NULL,
    CONSTRAINT "pk_identities" PRIMARY KEY ("id")
);

##########

CREATE INDEX "idx_identities_status" ON "identities" ("status");

