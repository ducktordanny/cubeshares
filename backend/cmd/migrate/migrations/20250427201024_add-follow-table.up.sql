CREATE TABLE IF NOT EXISTS "follow" (
  follower_id INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  following_id INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(follower_id, following_id)
);
