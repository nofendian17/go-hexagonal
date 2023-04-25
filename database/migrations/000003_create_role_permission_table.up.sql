CREATE TABLE IF NOT EXISTS role_permission (
    id              UUID NOT NULL PRIMARY KEY,
    permission_id   UUID NOT NULL,
    role_id         UUID NOT NULL,
    created_at      TIMESTAMP(0),
    updated_at      TIMESTAMP(0)
);