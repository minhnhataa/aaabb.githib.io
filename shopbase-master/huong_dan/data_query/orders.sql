use ecommerce;

create table orders(
id int primary key auto_increment,
user_id int not null,
total decimal not null,
payment_id int not null,
created_at timestamp,
modified_at timestamp,
deleted_at timestamp,

foreign key (payment_id) references user_payments(id),
foreign key (user_id) references users(id)
);

select*from orders;

delete from orders where id = 4;

