-- +migrate Up
-- +migrate StatementBegin

INSERT INTO ticket_statuses (id, name, created_at)
VALUES (1,
        'opn',
        1738378141
);

INSERT INTO ticket_statuses (id, name, created_at)
VALUES (2,
        'cld',
        1738378141
);

INSERT INTO ticket_statuses (id, name, created_at)
VALUES (3,
        'asn',
        1738378141
);

-- +migrate StatementEnd
-- +migrate Down
