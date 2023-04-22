create table if not exists roles
(
    id
    uuid
    not
    null
    primary
    key,
    name
    varchar
(
    255
) not null,
    active boolean default true not null,
    created_at timestamp
(
    0
),
    updated_at timestamp
(
    0
)
    );

