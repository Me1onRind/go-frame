create table `audio_tab` (
    `id` bigint(20) unsigned not null auto_increment,
    `filename` varchar(128) not null default '',
    `store_index` varchar(128) not null default '',
    `ctime` int(10) unsigned not null default 0,
    `mtime` int(10) unsigned not null default 0,
    primary key(`id`),
    uniq index uniq_idx_filename(`filename`)
) engine=Innodb default charset=utf8mb4;
