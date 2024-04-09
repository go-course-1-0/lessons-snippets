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
values ('Ivanovich', 'Ivanov', '+992987654321', '1995-02-27', 'ivan@mail.com', false, 'Ivan');

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

delete from students;

delete from students where first_name = 'hello';

delete from students where created_at < date(now());

truncate table students;