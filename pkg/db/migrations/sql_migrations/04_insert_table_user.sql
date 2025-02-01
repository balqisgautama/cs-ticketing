-- +migrate Up
-- +migrate StatementBegin

-- password sha256 dari password123
INSERT INTO users (client_id, email, password, is_active, created_at)
VALUES ('446e8caf-5da6-4fa9-a4ad-9df9ad1c88d0',
        'cs@info.sod.id',
        'ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f',
        1,
        1738378141
);

-- +migrate StatementEnd
-- +migrate Down
