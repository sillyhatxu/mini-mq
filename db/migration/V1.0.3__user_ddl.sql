CREATE TABLE IF NOT EXISTS user_info
(
  id                 INTEGER PRIMARY KEY AUTOINCREMENT,
  login_name         TEXT NOT NULL,
  password           TEXT NOT NULL,
  status             INTEGER  DEFAULT 1,
  created_time       datetime default current_timestamp,
  last_modified_time datetime default current_timestamp
);
CREATE INDEX idx_login_name ON user_info (login_name);
CREATE INDEX idx_status ON user_info (status);