CREATE TABLE `url_Store` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `short_url` varchar(200) DEFAULT NULL,
  `long_url` varchar(200) DEFAULT NULL,
  `expires_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
)