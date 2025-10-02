package posts

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ducktordanny/cubeshares/backend/db/queries"
	"github.com/ducktordanny/cubeshares/backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

var _ types.PostsStore = (*Store)(nil)

func (store *Store) ReadPostsOfUser(ctx context.Context, userId int64) ([]types.ReadPostResponseBody, error) {
	rows, err := store.db.QueryContext(ctx, queries.ReadPostsOfUserQuery, userId)
	if err != nil {
		return []types.ReadPostResponseBody{}, err
	}
	defer rows.Close()

	var posts []types.ReadPostResponseBody
	for rows.Next() {
		var b []byte
		if err := rows.Scan(&b); err != nil {
			return []types.ReadPostResponseBody{}, err
		}
		var post types.ReadPostResponseBody
		if err := json.Unmarshal(b, &post); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if posts == nil {
		return []types.ReadPostResponseBody{}, nil
	}
	return posts, nil
}

func (store *Store) CreateNewPost(ctx context.Context, userId int64, postRequestBody types.CreatePostRequestBody) (types.CreatePostResponseBody, error) {
	// TODO: implement later
	if postRequestBody.Average != nil {
	}
	// TODO: implement later
	if postRequestBody.Solve != nil {
	}

	var postId int64
	err := store.db.QueryRowContext(ctx, queries.CreateBasicPostQuery, userId, postRequestBody.Description).Scan(&postId)

	if err != nil {
		return types.CreatePostResponseBody{}, err
	}

	return types.CreatePostResponseBody{PostId: postId}, nil
}
