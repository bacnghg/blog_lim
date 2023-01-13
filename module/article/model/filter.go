package articlemodel

type FilterArticle struct {
	ArticleId   int    `json:"article_id" form:"article_id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	UsereId     int    `json:"user_id" form:"user_id"`
}
