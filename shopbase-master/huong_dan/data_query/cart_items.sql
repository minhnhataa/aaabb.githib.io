use ecommerce;

create table cart_items(
id int primary key auto_increment,
user_id int,
product_id varchar(20),
quantity int not null,
created_at timestamp,
modified_at timestamp,
foreign key (user_id) references users(id),
foreign key (product_id) references products(id)
);

select *from cart_items;

drop table people;

