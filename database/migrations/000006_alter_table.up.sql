ALTER TABLE
    user_role
ADD
    CONSTRAINT user_role_roles_id_foreign FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE;

ALTER TABLE
    user_role
ADD
    CONSTRAINT user_role_users_id_foreign FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE
    role_permission
ADD
    CONSTRAINT role_permission_role_id_foreign FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE;

ALTER TABLE
    role_permission
ADD
    CONSTRAINT role_permission_permission_id_foreign FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE;