CREATE TABLE IF NOT EXISTS "user" (
  id INT PRIMARY KEY,
  wca_id TEXT,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  avatar_url TEXT,
  role TEXT DEFAULT 'user',
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
