use ecommerce;

create table user_payments(
id int primary key auto_increment,
type varchar(20) not null,
user_id int,
foreign key (user_id) references users(id)
);

select *from user_payments

