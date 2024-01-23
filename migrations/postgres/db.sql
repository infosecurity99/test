create type user_role_enum as enum ('admin', 'customer');

create table users (
    id uuid primary key not null,
    full_name varchar(30),
    phone varchar(30) unique not null,
    password varchar(30) not null,
    user_role user_role_enum not null,
    cash int
);

create table baskets(
    id uuid primary key not null ,
    customer_id uuid references users(id) not null ,
    total_sum integer default 0
);
