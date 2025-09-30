package types

type PostsStore interface {
	CreateNewPost(userId int64, postRequestBody CreatePostRequestBody) (CreatePostResponseBody, error)
}

type CreatePostRequestBody struct {
	Solve       *SolveRequestBody   `json:"solve"`
	Average     *AverageRequestBody `json:"average"`
	Description string              `json:"description"`
}

type SolveRequestBody struct {
	SolveBase
	Category PuzzleCategory `json:"category"`
}

type AverageRequestBody struct {
	Solves   []SolveBase    `json:"solves"`
	Category PuzzleCategory `json:"category"`
	IsPR     *bool          `json:"isPR"`
}

type SolveBase struct {
	Penalty  *Penalty `json:"penalty"`
	Result   int64    `json:"result"`
	IsPR     *bool    `json:"isPR"`
	Scramble string   `json:"scramble"`
	Solution *string  `json:"solution"`
	Note     *string  `json:"note"`
}

type CreatePostResponseBody struct {
	PostId int64 `json:"postId"`
}
