create database avito_test;
create TABLE users(id serial primary key, first_name text, last_name text, active bool default true);
create table segment(id serial primary key, name text unique, active bool default true);
create table user_segment(id serial primary key, user_id integer references users(id), segment_id integer references segment(id), creation_time TIMESTAMP default now(), deletion_time T
 IMESTAMP, duration TIMESTAMP, active bool default true);
