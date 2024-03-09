create table users(
    user_id serial NOT NULL PRIMARY KEY,
    lastname varchar,
    firstname varchar,
    middlename varchar,
    email varchar,
    pswd varchar,
    passport varchar,
    inn varchar, 
    snils varchar,
    birthday varchar,
    role varchar
);

insert into users values ('adminov', 'admin', 'adminovich', 'admin@mail.ru', 'admin', '1234 567890', '12345678', '12345678', '2024-03-08', 'ADMIN');
