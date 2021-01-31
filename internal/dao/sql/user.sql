create table `user_tab` (
    `id` bigint(20) unsigned not null auto_increment,
    `username` varchar(64) not null default '',
    `password` char(16) not null default '',
    `ctime` int(10) unsigned not null default 0,
    `mtime` int(10) unsigned not null default 0,
    primary key(`id`)
) Engine=InnoDb default charset=utf8mb4;
