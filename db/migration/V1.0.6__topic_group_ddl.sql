CREATE TABLE IF NOT EXISTS topic_group
(
  topic_name         varchar(100) NOT NULL,
  topic_group        varchar(100) NOT NULL,
  offset             bigint(48)            DEFAULT 0,
  created_time       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  last_modified_time timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  PRIMARY KEY (topic_name, topic_group)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;