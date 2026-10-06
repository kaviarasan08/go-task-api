create table tasks (
id serial primary key,
title varchar(100) not null,
completed boolean not null default false
);

alter table tasks
add column created_at timestamp not null default current_timestamp;

alter table tasks
add column updated_at timestamp not null default current_timestamp;