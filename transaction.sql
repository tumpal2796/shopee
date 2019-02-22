create table transaction(
	billing_id SERIAL primary key,
	name varchar(125),
	tax_code smallint,
	price float
);