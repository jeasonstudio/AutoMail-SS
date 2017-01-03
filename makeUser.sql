USE jeason_daily;
CREATE TABLE IF NOT EXISTS `mail_ss` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `userid` varchar(40) DEFAULT NULL,
  `user_email` varchar(100) DEFAULT NULL,
  `is_receive` boolean DEFAULT TRUE,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;