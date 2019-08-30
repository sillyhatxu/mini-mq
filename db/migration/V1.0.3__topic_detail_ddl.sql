CREATE TABLE IF NOT EXISTS topic_detail
(
  topic              TEXT PRIMARY KEY,
  offset             INTEGER  DEFAULT 0,
  created_time       datetime default current_timestamp,
  last_modified_time datetime default current_timestamp
);
CREATE INDEX idx_topic ON topic_detail (topic);