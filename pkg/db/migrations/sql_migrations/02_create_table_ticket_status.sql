-- +migrate Up
-- +migrate StatementBegin

-- TICKET STATUS TABLE
CREATE SEQUENCE IF NOT EXISTS ticket_status_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "ticket_status" (
    "id" BIGINT DEFAULT nextval('ticket_status_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "name" VARCHAR(256) NOT NULL,
    "note" TEXT NULL,
    "created_at" INT NOT NULL,
    "updated_at" INT NULL,
    "deleted_at" INT NULL
);

-- +migrate StatementEnd
-- +migrate Down
