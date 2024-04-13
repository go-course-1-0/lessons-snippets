select *
from customer;

select *
from customer
order by customer_id;

select *
from customer
order by customer_id desc;

select *
from customer
order by customer_id desc
limit 10;

select *
from customer
order by customer_id desc
offset 10 limit 10;

select *
from customer
order by customer_id desc
offset 20 limit 10;

select *
from customer
order by customer_id desc
offset 30 limit 10;

select *
from customer
order by customer_id desc
offset 40 limit 10;



select *
from country;

select *
from country where country_id in (44,23,103);

select *
from city;

select country_id, count(*)
from city
group by country_id;

select country_id, count(*) as city_count from city group by country_id order by city_count desc;