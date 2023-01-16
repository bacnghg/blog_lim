package commentmodel

type FilterComment struct {
	CommentId   int    `json:"comment_id" form:"comment_id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	UsereId     int    `json:"user_id" form:"user_id"`
}
