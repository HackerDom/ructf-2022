create table if not exists public.users(
    id int generated always as identity,
    login varchar not null unique,
    created_at timestamp not null default now(),
    password_hash varchar not null,

    primary key (id)
);
