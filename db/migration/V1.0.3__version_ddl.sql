CREATE TABLE IF NOT EXISTS mini_mq_version
(
  id              INTEGER PRIMARY KEY AUTOINCREMENT,
  version         REAL NOT NULL,
  created_time    datetime default current_timestamp
);