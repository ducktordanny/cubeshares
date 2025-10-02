package types

import (
	"context"
	"time"
)

type PostsStore interface {
	ReadPostsOfUser(ctx context.Context, userId int64) ([]ReadPostResponseBody, error)
	CreateNewPost(ctx context.Context, userId int64, postRequestBody CreatePostRequestBody) (CreatePostResponseBody, error)
}

type ReadPostResponseBody struct {
	Id          int64               `json:"id"`
	UserId      int64               `json:"userId"`
	Solve       *SolveRequestBody   `json:"solve"`
	Average     *AverageRequestBody `json:"average"`
	Description string              `json:"description"`
	CreatedAt   time.Time           `json:"createdAt"`
}

type CreatePostRequestBody struct {
	Solve       *SolveRequestBody   `json:"solve"`
	Average     *AverageRequestBody `json:"average"`
	Description string              `json:"description"`
}

type SolveRequestBody struct {
	Penalty  *Penalty       `json:"penalty"`
	Result   int64          `json:"result"`
	IsPR     *bool          `json:"isPR"`
	Scramble string         `json:"scramble"`
	Solution *string        `json:"solution"`
	Note     *string        `json:"note"`
	Category PuzzleCategory `json:"category"`
}

type AverageRequestBody struct {
	Solves   []SolveRequestBody `json:"solves"`
	Category PuzzleCategory     `json:"category"`
	IsPR     *bool              `json:"isPR"`
}

type CreatePostResponseBody struct {
	PostId int64 `json:"postId"`
}
