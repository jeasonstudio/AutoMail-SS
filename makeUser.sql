USE jeason_daily;
CREATE TABLE IF NOT EXISTS `mail_ss` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `userid` varchar(40) DEFAULT NULL,
  `user_email` varchar(100) unique,
  `is_receive` boolean DEFAULT TRUE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

INSERT INTO `jeason_daily`.`mail_ss` (`id`, `name`, `userid`, `user_email`, `is_receive`) VALUES (NULL, '赵吉彤', '41524122', 'me@jeasonstudio.cn', '1')