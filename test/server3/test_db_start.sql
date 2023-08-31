create TABLE users(id serial primary key, first_name text, last_name text, active bool default true);

create table segment(id serial primary key, name text unique, active bool default true);

create table user_segment(id serial primary key, user_id integer references users(id), segment_name text references segment(name), creation_time BIGINT default EXTRACT(epoch from now()
 ), deletion_time BIGINT default 0, duration BIGINT default 0, active bool default true);

insert into users (first_name, last_name) values ('Person', 'Second'),('Krrr','Grrr'),('Jon','Sina'),('Miles','Green'),('Mason','Mount');
