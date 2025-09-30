package posts

import (
	"database/sql"
	"time"

	"github.com/ducktordanny/cubeshares/backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (store *Store) CreateNewPost(userId int64, postRequestBody types.CreatePostRequestBody) (types.CreatePostResponseBody, error) {
	createdAt := time.Now()

	// TODO: implement later
	if postRequestBody.Average != nil {
	}
	// TODO: implement later
	if postRequestBody.Solve != nil {
	}

	var postId int64
	err := store.db.QueryRow(`
		INSERT INTO "post" ("userId", "description", "createdAt")
		VALUES ($1, $2, $3)
		RETURNING id
	`, userId, postRequestBody.Description, createdAt).Scan(&postId)

	if err != nil {
		return types.CreatePostResponseBody{}, err
	}

	return types.CreatePostResponseBody{PostId: postId}, nil
}
