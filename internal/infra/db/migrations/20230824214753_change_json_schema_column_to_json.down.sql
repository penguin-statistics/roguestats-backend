SET statement_timeout = 0;

alter table public.researches alter column schema type jsonb using schema::jsonb;