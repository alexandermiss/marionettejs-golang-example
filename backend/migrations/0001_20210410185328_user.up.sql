create table if not exists users
(
    id bigserial not null primary key,
    name      varchar(20),
    user_name  varchar(15),
    email     varchar(30) not null,
    password  varchar(30) not null,
    role      varchar(15) not null,
    subscription varchar(10) not null,
    plan         varchar(10) not null,
    activated  boolean default false not null,
    verified   boolean default false not null,
    status     varchar(10) not null,
    is_private boolean default false not null
);


