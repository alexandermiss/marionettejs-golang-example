create table if not exists phrases
(
    id bigserial not null primary key,
    name          varchar(30),
    pronunciation varchar(60),
    language_id   integer not null,
    description   varchar(90) not null
);
