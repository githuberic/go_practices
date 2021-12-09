/*操作db*/
create database data_center default character set utf8mb4 collate utf8mb4_bin;

create table t_user (
                        `id` int(11) not null AUTO_INCREMENT primary key,
                        `name` VARCHAR(20) not null,
                        `mobile` VARCHAR(15) not null,
                        `address` VARCHAR(200) null
)ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;
