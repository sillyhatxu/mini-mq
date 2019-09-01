CREATE TABLE IF NOT EXISTS topic_group
(
  topic_name         TEXT NOT NULL,
  topic_group        TEXT NOT NULL,
  offset             INTEGER  DEFAULT 0,
  created_time       datetime default current_timestamp,
  last_modified_time datetime default current_timestamp,
  PRIMARY KEY (topic_name, topic_group)
);