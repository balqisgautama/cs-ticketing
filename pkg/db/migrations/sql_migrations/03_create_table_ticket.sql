-- +migrate Up
-- +migrate StatementBegin

-- TICKET TABLE
CREATE SEQUENCE IF NOT EXISTS ticket_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "tickets" (
    "id" BIGINT DEFAULT nextval('ticket_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "name" VARCHAR(256) NOT NULL,
    "msg" TEXT NULL,
    "user_id" BIGINT NOT NULL,
    "status_id" BIGINT NOT NULL,
    "created_at" INT NOT NULL,
    "updated_at" INT NULL,
    "deleted_at" INT NULL,
    CONSTRAINT "FK__users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    CONSTRAINT "FK__ticket_status" FOREIGN KEY ("status_id") REFERENCES "ticket_status" ("id") ON DELETE CASCADE
);

-- +migrate StatementEnd
-- +migrate Down
