
create table food_info(id serial primary key,
                       user_id int not null,
                       type int not null,
                       name text not null,
                       portion_size float not null);

insert into food_info values (1,1,1,'tea',200);