-- primary key ( = unique + not null)
-- foreign key
-- delete cascade and update cascade
-- check ( = if )
-- unique
-- not null
-- default

delete
from teachers
where id = 1;

-- will be foreign key violation error
delete
from teachers
where id = 2;

alter table courses
    drop constraint courses_teacher_id_fkey,
    add foreign key (teacher_id) references teachers (id) on delete cascade;

-- will not be foreign key violation error
delete
from teachers
where id = 2;

update teachers
set id = 10
where id = 3;

-- will be foreign key violation error
update teachers
set id = 10
where id = 4;

alter table courses
    drop constraint if exists courses_teacher_id_fkey,
    add foreign key (teacher_id) references teachers (id) on delete cascade on update cascade;

-- will not be foreign key violation error
update teachers
set id = 12
where id = 4;

alter table students
    add unique (phone),
    add unique (email);

alter table teachers
    add unique (phone),
    add unique (email);

insert into students (first_name, last_name, middle_name, phone)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987654444');

insert into students (first_name, last_name, middle_name, phone)
values ('Siyoif existsvush', 'Siyovushov', 'Siyovush', '+992987654321');

insert into students (first_name, last_name, middle_name, phone, email)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653333', 'ivan@mail.com');

insert into students (first_name, last_name, middle_name, phone, email)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653333', 'siyovush@mail.com');

insert into students (first_name, last_name, middle_name, phone, email, date_of_birth)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653334', 'siyovush@mail.com2', '1000-12-12');

alter table students
    add constraint check_age check ( date_part('year', now()) - date_part('year', date_of_birth) >= 16 and
                                     date_part('year', date_of_birth) >= 1900);

insert into students (first_name, last_name, middle_name, phone, email, date_of_birth)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653334', 'siyovush@mail.com2', '1000-12-12');

insert into students (first_name, last_name, middle_name, phone, email, date_of_birth)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653335', 'siyovush@mail.com3', '1900-12-12');

insert into students (first_name, last_name, middle_name, phone, email, date_of_birth)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653336', 'siyovush@mail.com4', '2020-12-12');

insert into students (first_name, last_name, middle_name, phone, email, date_of_birth)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653336', 'siyovush@mail.com4', '2009-12-12');

insert into students (first_name, last_name, middle_name, phone, email, date_of_birth)
values ('Siyovush', 'Siyovushov', 'Siyovush', '+992987653336', 'siyovush@mail.com4', '2008-12-12');

alter table teachers
    add column date_of_birth date default null;

alter table teachers
    add check ( date_part('year', now()) - date_part('year', date_of_birth) >= 16 and
                date_part('year', date_of_birth) >= 1900);

alter table teachers
    add check ( date_part('year', now()) - date_part('year', date_of_birth) >= 16 and
                date_part('year', date_of_birth) >= 1900);

alter table teachers
    add check ( date_part('year', now()) - date_part('year', date_of_birth) >= 16 and
                date_part('year', date_of_birth) >= 1900 and full_name = 'Siyovush');

create table if not exists students (id serial primary key );