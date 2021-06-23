create table if not exists users
(
    id SERIAL,
    first_name varchar(256),
    last_name varchar(256),
    nickname varchar(256),
    email varchar(256) not null,
    password varchar(256) not null,
    country varchar(256)
);
