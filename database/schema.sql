create table apps (
    id      int unsigned not null auto_increment,
    name    varchar(100),
    runtime varchar(20),

    primary key (`id`)
);