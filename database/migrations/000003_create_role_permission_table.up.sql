CREATE TABLE IF NOT EXISTS role_permission (
    permission_id   UUID NOT NULL,
    role_id         UUID NOT NULL,
    PRIMARY KEY (permission_id, role_id)
);