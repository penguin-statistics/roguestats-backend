SET statement_timeout = 0;

alter table public.researches alter column schema type json using schema::json;