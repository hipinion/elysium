CREATE TABLE `organzations` (
  `organization_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `organization_name` varchar(128) DEFAULT NULL,
  `organization_guid` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`organization_id`),
  KEY `organization_guid` (`organization_guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `publications` (
  `publication_id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `publication_name` varchar(128) NOT NULL DEFAULT '',
  `publication_guid` varchar(128) NOT NULL DEFAULT '',
  `organization_id` int(10) NOT NULL,
  PRIMARY KEY (`publication_id`),
  KEY `publication_guid` (`publication_guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
  `user_id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) NOT NULL DEFAULT '',
  `user_email` varchar(128) NOT NULL DEFAULT '',
  `user_first` varchar(32) DEFAULT NULL,
  `user_last` varchar(32) DEFAULT NULL,
  `user_phone` varchar(16) DEFAULT NULL,
  `user_reg_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;