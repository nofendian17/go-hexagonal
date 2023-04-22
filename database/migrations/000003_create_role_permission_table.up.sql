create table if not exists role_permission
(
    id
    uuid
    not
    null
    primary
    key,
    permission_id
    uuid
    not
    null,
    role_id
    uuid
    not
    null,
    created_at
    timestamp
(
    0
),
    updated_at timestamp
(
    0
)
    );

