create table students
(
    id            serial primary key not null,
    first_name    varchar(100),
    middle_name   varchar(100) default null,
    last_name     varchar(100)       not null,
    phone         varchar(30)        not null,
    email         varchar(100) default null,
    date_of_birth date         default null,
    is_active     boolean      default true,
    created_at    timestamptz  default now(),
    updated_at    timestamptz  default now(),
    deleted_at    timestamptz  default null
);

insert into students (first_name, middle_name, last_name, phone, email, date_of_birth, is_active)
values ('Ivan', 'Ivanovich', 'Ivanov', '+992987654321', 'ivan@mail.com', '1990-09-09', true);

insert into students (phone, email)
values ('+992987654321', 'ivan@mail.com');

insert into students (last_name, phone, email)
values ('Ivanov', '+992987654321', 'ivan@mail.com');

insert into students (last_name, phone)
values ('Ivanov', '+992987654321');

alter table students alter column first_name set not null;

alter table students add column address varchar(100);

alter table students drop column address;

alter table students add column address varchar(100) default 'Душанбе'
