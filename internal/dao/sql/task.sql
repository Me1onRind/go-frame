create table `task_tab` (
    `id` bigint(20) unsigned not null auto_increment,
    `task_name` varchar(64) not null default '',
    `args` varchar(2048) not null default '[]',
    `status` tinyint(3) unsigned not null default 0,
    `exec_times` int(10) unsigned not null default 0,
    `start_time` int(10) unsigned not null default 0,
    `last_exec_time` int(10) unsigned not null default 0,
    `next_exec_time` int(10) unsigned not null default 0,
    `exec_interval` int(10) unsigned not null default 0,
    `c_time` int(10) unsigned not null default 0,
    `m_time` int(10) unsigned not null default 0,
    primary key(`id`),
    index idx_next_exec_time(`next_exec_time`)
) Engine=InnoDB default charset=utf8mb4;
