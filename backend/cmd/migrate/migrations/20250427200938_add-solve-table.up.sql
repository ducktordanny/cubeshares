CREATE TABLE IF NOT EXISTS "solve" (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  category_id INT NOT NULL REFERENCES puzzle_category(id),
  penalty_id INT REFERENCES penalty(id),
  result INT NOT NULL,
  scramble TEXT NOT NULL,
  solution TEXT,
  note TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
