create table if not exists permissions (
    id uuid not null primary key,
    name varchar (255) not null,
    created_at timestamp (0),
    updated_at timestamp (0)
);