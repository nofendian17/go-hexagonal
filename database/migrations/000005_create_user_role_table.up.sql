CREATE TABLE IF NOT EXISTS user_role
(
    user_id    UUID NOT NULL,
    role_id    UUID NOT NULL,
    PRIMARY KEY (user_id, role_id)
);