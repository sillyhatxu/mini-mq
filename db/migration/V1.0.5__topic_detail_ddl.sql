CREATE TABLE IF NOT EXISTS topic_detail
(
  topic_name         varchar(100) NOT NULL PRIMARY KEY,
  offset             bigint(48)            DEFAULT 0,
  created_time       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  last_modified_time timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;