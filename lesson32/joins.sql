select *
from country;

select *
from city;

select *
from country
         join city on country.country_id = city.country_id;

select *
from country
         join city on country.country_id = city.country_id
order by country;

select *
from country
         join city on country.country_id = country_id
order by country;

select *
from country
         join city on country.country_id = city.country_id
order by country.country;

select *
from country
         join city c on country.country_id = c.country_id
order by country.country;

select country.country_id, country.country, country.last_update, count(c.city_id)
from country
         join city c on country.country_id = c.country_id
group by country.country_id
order by country.country;

select country.country_id, country.country, country.last_update, count(c.city_id)
from country
         join city c on country.country_id = c.country_id
group by country.country_id
order by count desc;

select country.country_id, country.country, country.last_update, count(c.city_id)
from country
         join city c on country.country_id = c.country_id
group by country.country_id
order by count desc
limit 10;

select country.country_id, country.country, count(c.city_id) as "Количество городов"
from country
         join city c on country.country_id = c.country_id
group by country.country_id
order by "Количество городов" desc
limit 10;

-- last query for getting top countries
select country.country_id, country.country, count(c.city_id) as number_of_cities
from country
         join city c on country.country_id = c.country_id
group by country.country_id
order by number_of_cities desc
limit 10;



select country.country_id, country.country, count(c.city_id) as number_of_cities
from country
         left join city c on country.country_id = c.country_id
group by country.country_id
order by number_of_cities desc;

select country.country_id, country.country, count(c.city_id) as number_of_cities
from country
         left outer join city c on country.country_id = c.country_id
group by country.country_id
order by number_of_cities desc;


select *
from country
where (select count(*) from city where city.country_id = country.country_id) >= 30
order by country.country;

select *
from country
where country_id > 80;

select *
from country
         join city c on country.country_id = c.country_id
         join address a on c.city_id = a.city_id
         join customer c2 on a.address_id = c2.address_id;

select country.country_id,
       country.country,
       count(c.city_id)    as number_of_cities,
       count(a.address_id) as number_of_addresses
from country
         left join city c on country.country_id = c.country_id
         left join address a on c.city_id = a.city_id
group by country.country_id
order by number_of_cities desc;

select city.city_id, city.city, count(a.address_id) as number_of_addresses
from city
         join address a on city.city_id = a.city_id
group by city.city_id
order by number_of_addresses desc;

select country.country_id,
       country.country,
       count(c.city_id)    as number_of_cities,
       count(a.address_id) as number_of_addresses
from country
         join city c on country.country_id = c.country_id
         join address a on c.city_id = a.city_id
where country.country_id in (select city.country_id
                             from city
                             where (select count(*) from address where city.city_id = address.city_id) >= 2)
group by country.country_id
order by country.country;

select city.country_id, city.city, address.address
from city
         join address on address.city_id = city.city_id
where (select count(*) from address where city.city_id = address.city_id) >= 2
