use ecommerce;

create table order_details(
id int primary key auto_increment,
product_id varchar(20) not null,
quantity int not null,
order_id int not null,
created_at timestamp,
modified_at timestamp,
deleted_at timestamp,

foreign key (product_id) references products(id),
foreign key (order_id) references orders(id)
);

select *from order_details;