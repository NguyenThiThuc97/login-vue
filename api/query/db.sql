create table Users (
	id serial,
	username varchar,
	email varchar,
	password varchar,
	token varchar,
	token_verified_at timestamp
);

create table Products (
	id serial,
	product_name varchar,
	del_fg integer
)
