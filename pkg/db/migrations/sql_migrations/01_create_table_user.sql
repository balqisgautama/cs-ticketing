-- +migrate Up
-- +migrate StatementBegin

-- USERS TABLE
CREATE SEQUENCE IF NOT EXISTS user_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "users" (
    "id" BIGINT DEFAULT nextval('user_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "client_id" VARCHAR(256) NOT NULL UNIQUE,
    "email" VARCHAR(256) NOT NULL UNIQUE,
    "password" VARCHAR(256) NOT NULL,
    "phone" VARCHAR(256) NULL UNIQUE,
    "is_active" INT DEFAULT 0,
    "created_at" INT NOT NULL,
    "updated_at" INT NULL,
    "deleted_at" INT NULL
);

-- +migrate StatementEnd
-- +migrate Down
