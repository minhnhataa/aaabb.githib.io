use ecommerce;

create table users(
id int primary key auto_increment,
name varchar(45) not null,
email varchar(45) unique not null,
password varchar(45) not null
);

select* from users;