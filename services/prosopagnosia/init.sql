create table if not exists public.demos(
    name varchar not null unique,
    author varchar not null,
    secret varchar not null,
    key varchar not null,
    rom_path varchar not null,
    created_at timestamp not null default now(),

    primary key (name)
);

