
create table address
(
    id          int auto_increment
        primary key,
    street_name varchar(100) null,
    zip_code    varchar(12)  null,
    city        varchar(100) null,
    `state`       varchar(60)  null,
    country     varchar(30)  null
);

create table person
(
    id         int auto_increment
        primary key,
    first_name varchar(100) null,
    last_name  varchar(100) null,
    age        int          null,
    address_id int          null,
    constraint person_address_id_fk
        foreign key (address_id) references address (id)
            on delete cascade
);