CREATE TABLE IF NOT EXISTS "puzzle_category" (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

INSERT INTO "puzzle_category" (name)
VALUES
  ('3x3'),
  ('2x2'),
  ('4x4'),
  ('5x5'),
  ('6x6'),
  ('7x7'),
  ('3x3oh'),
  ('3x3bl'),
  ('3x3mbl'),
  ('4x4bl'),
  ('5x5bl'),
  ('Megaminx'),
  ('Pyraminx'),
  ('Skewb'),
  ('Square-1'),
  ('Clock'),
  ('FMC');
