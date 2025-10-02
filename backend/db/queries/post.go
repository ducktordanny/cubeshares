package queries

var ReadPostsOfUserQuery string = baseReadPostQuery + `
	WHERE p."userId" = $1
	ORDER BY p."createdAt" DESC;
`

var CreateBasicPostQuery string = `
	INSERT INTO "post" ("userId", "description")
	VALUES ($1, $2)
	RETURNING id
`

var baseReadPostQuery string = `
	SELECT
		jsonb_build_object(
			'id',          p.id,
			'userId',      p."userId",
			'description', p.description,
			'createdAt',   p."createdAt",
			'solve',
				CASE WHEN s.id IS NULL THEN NULL ELSE
					jsonb_build_object(
						'category', pc.name,
						'penalty',  pen.type,
						'result',   s.result,
						'isPR',     s."isPR",
						'scramble', s.scramble,
						'solution', s.solution,
						'note',     s.note
					)
				END,
			'average',
				CASE WHEN a.id IS NULL THEN NULL ELSE
					jsonb_build_object(
						'isPR',     a."isPR",
						'note',     a.note,
						'solves',
							(
								SELECT jsonb_agg(
												 jsonb_build_object(
													 'penalty',  pen2.type,
													 'result',   s2.result,
													 'isPR',     s2."isPR",
													 'scramble', s2.scramble,
													 'solution', s2.solution,
													 'note',     s2.note
												 )
											 )
								FROM "averageSolve" asv
								JOIN solve s2                  ON s2.id   = asv."solveId"
								LEFT JOIN "puzzleCategory" pc2 ON pc2.id  = s2."categoryId"
								LEFT JOIN penalty pen2         ON pen2.id = s2."penaltyId"
								WHERE asv."averageId" = a.id
							)
					)
				END
		) AS post_json
	FROM post p
	LEFT JOIN solve s               ON s.id   = p."solveId"
	LEFT JOIN average a             ON a.id   = p."averageId"
	LEFT JOIN "puzzleCategory" pc   ON pc.id  = s."categoryId"
	LEFT JOIN penalty pen           ON pen.id = s."penaltyId"
`
