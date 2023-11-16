CREATE TABLE people ( 
last_name varchar(50) not null, 
first_name varchar(50) not null, 
dob date not null, 
gender enum('m', 'f')not null, 
key(last_name, first_name, dob) 
);

insert into people (last_name, first_name, dob, gender)
values("Trong", "Tran", '2000-12-12', "f"); 

select *from people;
explain select * from people where first_name='S%' and last_name='Peter';
explain select * from people where dob='1990-10-05' and first_name='Smitth';
explain select * from people where last_name='Peter' and dob='1990-10-05' and first_name='Smitth';
explain select * from people where last_name='Trong'; 
explain select * from people where last_name like 'P%';
explain select * from people where last_name='Peter' and first_name like 'S%';
explain select * from people where last_name like  '%r';
explain select * from people where last_name='Peter' and dob='1990-10-05';