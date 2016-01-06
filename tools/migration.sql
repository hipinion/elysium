CREATE TABLE `organzations` (
  `organization_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `organization_name` varchar(128) DEFAULT NULL,
  `organization_guid` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`organization_id`),
  KEY `organization_guid` (`organization_guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;