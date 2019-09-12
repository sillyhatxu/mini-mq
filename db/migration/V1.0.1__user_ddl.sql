CREATE TABLE IF NOT EXISTS user_info
(
  id                 bigint(48)                                NOT NULL AUTO_INCREMENT PRIMARY KEY,
  login_name         varchar(100)                              NOT NULL,
  password           varchar(100)                              NOT NULL,
  status             int          DEFAULT 1,
  created_time       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  last_modified_time timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  INDEX (login_name, status)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;