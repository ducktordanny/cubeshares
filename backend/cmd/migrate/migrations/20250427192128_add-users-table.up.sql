CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY,
  wca_id TEXT,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  gender TEXT NOT NULL,
  bio TEXT DEFAULT '',
  country_iso TEXT NOT NULL,
  avatar_url TEXT,
  role TEXT DEFAULT 'user',
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
