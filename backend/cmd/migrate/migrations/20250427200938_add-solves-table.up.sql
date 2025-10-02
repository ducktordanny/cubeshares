CREATE TABLE IF NOT EXISTS solves (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  category_id INT NOT NULL REFERENCES puzzle_categories(id),
  result INT NOT NULL,
  is_pr BOOLEAN NOT NULL DEFAULT FALSE,
  penalty TEXT CHECK (penalty IN ('+2', 'DNF') OR penalty IS NULL),
  scramble TEXT NOT NULL,
  solution TEXT,
  note TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
