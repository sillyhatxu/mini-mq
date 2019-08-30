CREATE TABLE IF NOT EXISTS topic_group
(
  id                 INTEGER PRIMARY KEY AUTOINCREMENT,
  topic              TEXT NOT NULL,
  topic_group        TEXT NOT NULL,
  offset             INTEGER  DEFAULT 0,
  created_time       datetime default current_timestamp,
  last_modified_time datetime default current_timestamp
);
CREATE INDEX idx_topic_name_and_group ON topic_group (topic, topic_group);