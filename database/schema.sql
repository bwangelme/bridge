create table apps
(
    id      int unsigned not null auto_increment,
    name    varchar(100),
    runtime varchar(20),
    image   varchar(100),

    UNIQUE KEY `uni_name` (`name`),
    primary key (`id`)
);