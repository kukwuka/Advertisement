CREATE TABLE   advertisement
(
    id serial not null unique,
    date date not null,
    name varchar(200),
    description varchar(1000),
    cost   float not null,
    img_url varchar(200)[]
);