-- +migrate Up
-- +migrate StatementBegin

-- password sha256 dari password123
INSERT INTO users (id, client_id, email, password, is_active, created_at)
VALUES (1,
        '446e8caf-5da6-4fa9-a4ad-9df9ad1c88d0',
        'cs01@info.sod.id',
        'ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f',
        1,
        1738378141
);

-- password sha256 dari password123
INSERT INTO users (id, client_id, email, password, is_active, created_at)
VALUES (2,
        '507fb9a5-6e30-41fb-91e5-324fa8cb6d30',
        'cs02@info.sod.id',
        'ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f',
        1,
        1738378141
);

-- password sha256 dari password123
INSERT INTO users (id, client_id, email, password, is_active, created_at)
VALUES (3,
        '392ebedb-1440-4e73-b2f7-a88c4af5528f',
        'cs03@info.sod.id',
        'ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f',
        1,
        1738378141
);

-- +migrate StatementEnd
-- +migrate Down
