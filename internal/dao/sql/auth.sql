create table `auth_tab` (
    `id` bigint(20) unsigned not null auto_increment,
    `app_name` varchar(64) not null default '',
    `app_secret` char(16) not null default '',
    `config_id` bigint(20) unsigned not null default 0,
    `ctime` int(10) unsigned not null default 0,
    `mtime` int(10) unsigned not null default 0,
    primary key(`id`),
    unique key uniq_idx_app_name(`app_name`)
) engine=Innodb default charset=utf8mb4;

create table `auth_config_tab` (
    `id` bigint(20) unsigned not null auto_increment,
    `config_name` varchar(32) not null default '',
    `expires` int(10) unsigned not null default 0,
    `flag` bigint(20) unsigned not null default 0,
    `ctime` int(10) unsigned not null default 0,
    `mtime` int(10) unsigned not null default 0,
    primary key(`id`)
) engine=Innodb default charset=utf8mb4;



