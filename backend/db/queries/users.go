package queries

var ReadUserByIdQuery string = `SELECT * FROM users WHERE id = $1`

var UpdateUserBioQuery string = `UPDATE users SET bio = $1 WHERE id = $2`

var InsertOrUpdateUserQuery string = `
	INSERT INTO users (id, wca_id, name, email, gender, country_iso, avatar_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (id) DO UPDATE
	SET wca_id = EXCLUDED.wca_id,
			name = EXCLUDED.name,
			email = EXCLUDED.email,
			avatar_url = EXCLUDED.avatar_url
`
