CREATE TABLE IF NOT EXISTS "post" (
  "id" SERIAL PRIMARY KEY,
  "userId" INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  "averageId" INT REFERENCES average(id),
  "solveId" INT REFERENCES solve(id),
  "description" TEXT,
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW()
);
