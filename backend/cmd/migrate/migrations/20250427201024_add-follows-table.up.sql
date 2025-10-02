CREATE TABLE IF NOT EXISTS follows (
  follower_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  followed_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY(follower_id, followed_id)
);
