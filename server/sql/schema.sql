drop table if exists order_products;
drop table if exists "orders";
drop table if exists products;
drop table if exists categories;

create table categories (
    id serial primary key,
    name varchar(255) not null
);

create table products(
    id serial primary key,
    label varchar(255) not null,
    type int references categories(id) not null,
    downloadURL varchar(255),
    weight float,
    created_at timestamp default current_timestamp
);

create table "orders"(
    id serial primary key,
    user_id int not null,
    created_at timestamp default current_timestamp
);

create table order_products(
    id serial primary key,
    order_id int references "orders"(id) not null,
    product_id int references products(id) not null,
    quantity int not null,
    created_at timestamp default current_timestamp
);