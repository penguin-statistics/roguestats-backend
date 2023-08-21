create table
    public.users (
        user_id text not null constraint users_pk primary key,
        name text not null,
        email text not null,
        credential text not null,
        attributes jsonb
    );

create unique index users_email_uindex on public.users (email);

alter table public.users add constraint users_email unique (email);

create table
    researches (
        research_id text not null constraint research_pk primary key,
        name text not null,
        schema jsonb not null
    );

create table
    events (
        event_id text not null constraint events_pk primary key,
        created_at timestamp,
        user_id text constraint events_users_user_id_fk references users,
        user_agent text,
        content jsonb,
        research_id text constraint events_researches_research_id_fk references researches
    );