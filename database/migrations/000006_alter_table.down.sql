ALTER TABLE
    user_role
    DROP CONSTRAINT IF EXISTS user_role_roles_id_foreign;
ALTER TABLE
    user_role
    DROP CONSTRAINT IF EXISTS user_role_users_id_foreign;

ALTER TABLE
    role_permission
    DROP CONSTRAINT IF EXISTS role_permission_role_id_foreign;

ALTER TABLE
    role_permission
    DROP CONSTRAINT IF EXISTS role_permission_permission_id_foreign;