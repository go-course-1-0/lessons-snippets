create database academy_db;

create table if not exists admins
(
    id         serial primary key,
    full_name  varchar(100)        not null,
    email      varchar(100) unique not null,
    password   varchar(100)        not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists teachers
(
    id         serial primary key,
    full_name  varchar(100) not null,
    subject    varchar(100) not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists courses
(
    id         serial primary key,
    title      varchar(100) not null,
    duration   integer      not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists groups
(
    id         serial primary key,
    course_id  integer references courses (id)  not null,
    teacher_id integer references teachers (id) not null,
    title      varchar(100)                     not null,
    start      date        default null,
    finish     date        default null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists students
(
    id            serial primary key,
    group_id      integer references groups (id) not null,
    full_name     varchar(100)                   not null,
    phone         varchar(20) unique             not null,
    date_of_birth date                           not null,
    created_at    timestamptz default now(),
    updated_at    timestamptz default now()
);

create table if not exists lessons
(
    id          serial primary key,
    group_id    integer references groups (id) not null,
    day_of_week integer                        not null,
    time        time        default null,
    created_at  timestamptz default now(),
    updated_at  timestamptz default now()
);

