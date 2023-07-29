create table
    public.users (
        user_id serial not null constraint users_pk primary key,
        name text not null,
        email text not null,
        credential text not null,
        attributes jsonb
    );

create unique index users_email_uindex on public.users (email);

alter table public.users add constraint users_email unique (email);