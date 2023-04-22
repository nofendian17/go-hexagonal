ALTER TABLE users DROP CONSTRAINT IF EXISTS users_role_id_foreign;
ALTER TABLE role_permission DROP CONSTRAINT IF EXISTS role_permission_role_id_foreign;
ALTER TABLE role_permission DROP CONSTRAINT IF EXISTS role_permission_permission_id_foreign;
