create TABLE users(id serial primary key, first_name text, last_name text, active bool default true);

--insert into users(id, first_name, last_name) values (3, "David", "Brown");

create table segment(id serial primary key, name text unique, active bool default true);

--insert into segment(name) values ('AVITO_TEST'),('AVITO_MESSAGE'),('OTIVA_VOICE'),('OTIVA_TEST'),('AVITO_VOICE'),('AVITO_WRITE');

create table user_segment(id serial primary key, user_id integer references users(id), segment_name text references segment(name), creation_time BIGINT default EXTRACT(epoch from now()
 ), deletion_time BIGINT, duration BIGINT, active bool default true);

--insert into user_segment(user_id, segment_name) values ()
