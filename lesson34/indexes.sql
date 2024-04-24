-- pros:
-- faster data querying
-- less operations need to be done

-- cons:
-- storage
-- build costs time and operations
-- build blocks insert, update and delete queries

select * from address;

select * from address where phone='195003555232';

explain select * from address;

explain select * from address where phone='195003555232';

create index on address(phone);

explain select * from address;

explain select * from address where phone='195003555232';

select count(*) from payment;

select * from payment where amount = 2.99;

explain select * from payment where amount = 2.99;

explain analyze select * from payment where amount = 2.99;

create index on payment(amount);

explain select * from payment where amount = 2.99;

explain analyze select * from payment where amount = 2.99;
