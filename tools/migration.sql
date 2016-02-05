
CREATE TABLE `forums` (
  `forum_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `forum_guid` varchar(256) DEFAULT NULL,
  `forum_name` varchar(256) DEFAULT NULL,
  `organization_id` int(10) DEFAULT NULL,
  PRIMARY KEY (`forum_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table hosts
# ------------------------------------------------------------

CREATE TABLE `hosts` (
  `host_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `host_guid` varchar(256) DEFAULT NULL,
  `organization_id` int(10) DEFAULT NULL,
  PRIMARY KEY (`host_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table organzations
# ------------------------------------------------------------

CREATE TABLE `organzations` (
  `organization_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `organization_name` varchar(128) DEFAULT NULL,
  `organization_guid` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`organization_id`),
  KEY `organization_guid` (`organization_guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table posts
# ------------------------------------------------------------

CREATE TABLE `posts` (
  `post_id` int(14) unsigned NOT NULL AUTO_INCREMENT,
  `topic_id` int(12) DEFAULT NULL,
  `post_guid` varchar(256) DEFAULT NULL,
  `post_title` varchar(256) DEFAULT NULL,
  `post_text` mediumtext,
  `user_id` int(12) DEFAULT NULL,
  PRIMARY KEY (`post_id`),
  KEY `topic_id` (`topic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table publications
# ------------------------------------------------------------

CREATE TABLE `publications` (
  `publication_id` int(9) unsigned NOT NULL AUTO_INCREMENT,
  `publication_name` varchar(128) NOT NULL DEFAULT '',
  `publication_guid` varchar(128) NOT NULL DEFAULT '',
  `organization_id` int(10) NOT NULL,
  PRIMARY KEY (`publication_id`),
  KEY `publication_guid` (`publication_guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table sessions
# ------------------------------------------------------------

CREATE TABLE `sessions` (
  `session_id` varchar(128) NOT NULL DEFAULT '',
  `user_id` int(12) DEFAULT NULL,
  `session_created` timestamp NULL DEFAULT NULL,
  `session_updated` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `session_ip` varchar(128) DEFAULT NULL,
  `session_user_agent` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`session_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table topics
# ------------------------------------------------------------

CREATE TABLE `topics` (
  `topic_id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `forum_id` int(10) DEFAULT NULL,
  `topic_guid` varchar(256) DEFAULT NULL,
  `topic_title` varchar(256) DEFAULT '',
  `user_id` int(12) NOT NULL,
  PRIMARY KEY (`topic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table users
# ------------------------------------------------------------

CREATE TABLE `users` (
  `user_id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) NOT NULL DEFAULT '',
  `user_email` varchar(128) NOT NULL DEFAULT '',
  `user_first` varchar(32) DEFAULT NULL,
  `user_last` varchar(32) DEFAULT NULL,
  `user_phone` varchar(16) DEFAULT NULL,
  `user_reg_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_salt` varchar(256) DEFAULT NULL,
  `user_password` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


