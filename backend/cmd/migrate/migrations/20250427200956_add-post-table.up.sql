CREATE TABLE IF NOT EXISTS "post" (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  average_id INT REFERENCES average(id),
  solve_id INT REFERENCES solve(id),
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
