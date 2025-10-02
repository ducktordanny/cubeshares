package queries

var ReadPostsOfUserQuery string = baseReadPostQuery + `
	WHERE p.user_id = $1
	ORDER BY p.created_at DESC;
`

var CreateBasicPostQuery string = `
	INSERT INTO posts (user_id, description)
	VALUES ($1, $2)
	RETURNING id
`

var baseReadPostQuery string = `
	SELECT
		jsonb_build_object(
			'id',          p.id,
			'userId',      p.user_id,
			'description', p.description,
			'createdAt',   p.created_at,
			'solve',
				CASE WHEN s.id IS NULL THEN NULL ELSE
					jsonb_build_object(
						'category', pc.name,
						'penalty',  s.penalty,
						'result',   s.result,
						'isPR',     s.is_pr,
						'scramble', s.scramble,
						'solution', s.solution,
						'note',     s.note
					)
				END,
			'average',
				CASE WHEN a.id IS NULL THEN NULL ELSE
					jsonb_build_object(
						'isPR',     a.is_pr,
						'note',     a.note,
						'solves',
							(
								SELECT jsonb_agg(
												 jsonb_build_object(
													 'penalty',  s2.penalty,
													 'result',   s2.result,
													 'isPR',     s2.is_pr,
													 'scramble', s2.scramble,
													 'solution', s2.solution,
													 'note',     s2.note
												 )
											 )
								FROM average_solves asv
								JOIN solves s2                  ON s2.id   = asv.solve_id
								LEFT JOIN puzzle_categories pc2 ON pc2.id  = s2.category_id
								WHERE asv.average_id = a.id
							)
					)
				END
		) AS post_json
	FROM posts p
	LEFT JOIN solves s               ON s.id   = p.solve_id
	LEFT JOIN averages a             ON a.id   = p.average_id
	LEFT JOIN puzzle_categories pc   ON pc.id  = s.category_id
`
