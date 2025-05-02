CREATE TABLE IF NOT EXISTS "penalty" (
  "id" SERIAL PRIMARY KEY,
  "type" TEXT NOT NULL CHECK (type IN ('+2', 'DNF'))
);

INSERT INTO "penalty" (type)
VALUES
  ('+2'),
  ('DNF');
