CREATE TABLE IF NOT EXISTS mini_mq_version
(
  id           bigint(48)   NOT NULL AUTO_INCREMENT PRIMARY KEY,
  version      float        NOT NULL,
  created_time timestamp(3) NOT NULL DEFAULT current_timestamp(3)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;