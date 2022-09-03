CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table users (id text, store_id text, name text, username text, password text, role text);
insert into users values ('e438d782-24c6-4d6d-bd02-2a4081160217', '1', 'Michael', 'admin', 'admin', 'admin');