-- select - для выборки (получения) данных из таблицы
-- insert - для ввода данных в таблицу
-- update - для изменения (обновления) записи в таблице
-- delete - для удаления записи из таблицы

-- where - для уточнения запроса
-- order by - для сортировки результата

select *
from students;

select id, first_name, phone
from students;

insert into students (middle_name, last_name, phone, date_of_birth, email, is_active, first_name)
values ('Ivanovich', 'Ivanov 4', '+992987654321', '1995-02-27', 'ivan@mail.com', false, 'Ivan 123');

insert into students (middle_name, last_name, phone, date_of_birth, email, is_active, first_name)
values ('Ivanovich 1', 'Ivanov 8', '+992987654321', '1995-02-27', 'ivan@mail.com', false, 'Ivan 25');

insert into students (middle_name, last_name, phone, date_of_birth, email, is_active, first_name)
values ('Ivanovich 2', 'Ivanov 9', '+992987654321', '1995-02-27', 'ivan@mail.com', false, 'Ivan 67');

insert into students (middle_name, last_name, phone, date_of_birth, email, is_active, first_name)
values ('Ivanovich 3', 'Ivanov 12', '+992987654321', '1995-02-27', 'ivan@mail.com', false, 'Ivan 90');

insert into students default
values;

select *
from students;

select *
from students
where is_active = true;

select *
from students
where is_active = true
  and middle_name is not null;

select *
from students
where is_active = true
  and (middle_name is not null or first_name = 'hello');

select *
from students
where id <= 4;

select *
from students
where first_name = 'Ivan';

select *
from students
where first_name like '%3';

update students
set is_active = false;

update students
set is_active = true
where id = 13
   or id = 14;

update students
set is_active= true
where id > 5;

update students
set first_name = 'Siyovush'
where id = 3;

delete
from students;

delete
from students
where first_name = 'hello';

delete
from students
where created_at < date(now());

truncate table students;

select *
from students
where id in (23, 25, 27, 50, 78);

delete
from students
where id in (23, 25, 27, 50, 78);

-- default is ascending order
select * from students order by first_name;
select * from students order by middle_name;
select * from students order by last_name;
select * from students order by id;

-- yavno ascending
select * from students order by first_name asc;
select * from students order by middle_name asc;
select * from students order by last_name asc;
select * from students order by id asc;

-- yavno descending
select * from students order by first_name desc;
select * from students order by middle_name desc;
select * from students order by last_name desc;
select * from students order by id desc;


select * from students limit 10;

select count(*) from students;

select count(*), is_active from students group by is_active;

select count(first_name), is_active from students group by is_active;

select sum(id), count(id), is_active from students group by is_active;
