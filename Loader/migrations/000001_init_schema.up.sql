create table files
(
    id         varchar not null
        constraint files_pk
            primary key,
    "Name"     varchar,
    "dtCreate" date,
    "Path"     varchar,
    "Size"     integer,
    "Author"   varchar
);
