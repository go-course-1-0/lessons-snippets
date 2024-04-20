create table teachers
(
    id         serial primary key not null,
    full_name  varchar(200),
    phone      varchar(30)        not null,
    email      varchar(100) default null,
    created_at timestamptz  default now(),
    updated_at timestamptz  default now(),
    deleted_at timestamptz  default null
);

insert into teachers (full_name, phone, email)
values ('Ivan Ivanovich Ivanov', '+992987654321', 'ivan@mail.com');

insert into teachers (full_name, phone, email)
values ('Ivan Ivanovich Ivanov 1', '+992987654321', 'ivan@mail.com');

insert into teachers (full_name, phone, email)
values ('Ivan Ivanovich Ivanov 2', '+992987654321', 'ivan@mail.com');

insert into teachers (full_name, phone, email)
values ('Ivan Ivanovich Ivanov 3', '+992987654321', 'ivan@mail.com');


create table courses
(
    id         serial primary key not null,
    teacher_id integer,
    title      varchar(100),
    duration   integer,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz default null,

    constraint courses_teacher_id_fkey
        foreign key (teacher_id)
            REFERENCES teachers (id)
);

insert into courses (teacher_id, title, duration)
values (4, 'Go', 4);

insert into courses (teacher_id, title, duration)
values (4, 'DevOps', 4);

insert into courses (teacher_id, title, duration)
values (2, 'React', 3);

insert into courses (teacher_id, title, duration)
values (2, 'Vue', 3);

insert into courses (teacher_id, title, duration)
values (40, 'Golang', 4);


select *
from teachers;

select *
from teachers
         left join courses c on teachers.id = c.teacher_id
order by teachers.id;

select teachers.id, teachers.full_name, count(*) as courses_count
from teachers
         left join courses c on teachers.id = c.teacher_id
group by teachers.id
order by teachers.id;

select teachers.id, teachers.full_name, count(*) as courses_count
from teachers
         join courses c on teachers.id = c.teacher_id
group by teachers.id
order by teachers.id;
