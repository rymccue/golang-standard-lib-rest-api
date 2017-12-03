create table users (
    id serial primary key,
    name varchar(60) not null,
    email varchar(150) not null,
    password char(64) not null,
    salt char(32) not null,
    created_at timestamp default current_timestamp
);

create table jobs (
    id serial primary key,
    title varchar(150) not null,
    description text not null,
    user_id int not null,
    created_at timestamp default current_timestamp
);
