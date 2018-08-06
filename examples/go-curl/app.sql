CREATE TABLE `t_user_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `salt` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `is_current` int(11) NOT NULL DEFAULT '0',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`,`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `t_user_base` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login_name` varchar(45) NOT NULL,
  `name` varchar(45) NOT NULL DEFAULT '',
  `idc` varchar(45) NOT NULL DEFAULT '',
  `nickname` varchar(45) NOT NULL,
  `age` int(11) NOT NULL DEFAULT '0',
  `male` int(11) NOT NULL DEFAULT '0',
  `avatar` varchar(45) NOT NULL DEFAULT '',
  `location` varchar(45) NOT NULL DEFAULT 'shenzhen',
  `profession` varchar(45) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `token` varchar(45) NOT NULL DEFAULT '',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;