create table if not exists profiles
(
    id         bigint auto_increment,
    email      varchar(255)  not null,
    password   varchar(1000) not null,
    first_name varchar(255)  not null,
    last_name  varchar(255)  not null,
    city       varchar(255)  not null,
    age        int           not null,
    interests  text          null,
    constraint profiles_pk
        primary key (id)
);

# create unique index udx_profiles_email
#     on profiles (email);