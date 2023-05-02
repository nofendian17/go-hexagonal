create table if not exists users (
    id uuid not null primary key,
    name varchar (255) not null,
    email varchar (255) not null constraint users_email_unique unique,
    active boolean default true not null,
    email_verified_at timestamp (0),
    salt varchar (255) not null,
    password varchar (255) not null,
    remember_token varchar (100),
    created_at timestamp (0),
    updated_at timestamp (0)
);