package commentmodel

type Filter struct {
	ArticleId   int    `json:"articles_id" form:"articles_id"`
	UserId      int    `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
}
