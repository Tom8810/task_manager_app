ALTER TABLE projects
  MODIFY COLUMN total_time BIGINT DEFAULT 0 NOT NULL,
  MODIFY COLUMN goal_time BIGINT NOT NULL;