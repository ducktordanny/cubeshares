CREATE TABLE IF NOT EXISTS "follow" (
  "followerId" INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  "followingId" INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
  "createdAt" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY("followerId", "followingId")
);
