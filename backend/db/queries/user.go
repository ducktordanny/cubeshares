package queries

var ReadUserByIdQuery string = `SELECT * FROM "user" WHERE "id" = $1`

var UpdateUserBioQuery string = `UPDATE "user" SET "bio" = $1 WHERE "id" = $2`

var InsertOrUpdateUserQuery string = `
	INSERT INTO "user" ("id", "wcaId", "name", "email", "gender", "countryISO", "avatarURL")
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT ("id") DO UPDATE
	SET "wcaId" = EXCLUDED."wcaId",
			"name" = EXCLUDED."name",
			"email" = EXCLUDED."email",
			"avatarURL" = EXCLUDED."avatarURL"
`
