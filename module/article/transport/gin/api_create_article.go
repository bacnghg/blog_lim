package articlegin

import (
	"gin_form/common"
	articlebusiness "gin_form/module/article/business"
	articlemodel "gin_form/module/article/model"
	articlestorage "gin_form/module/article/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data articlemodel.ArticleCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := articlestorage.NewDbStore(db)
		business := articlebusiness.NewCreateArticle(storage)

		if err := business.CreateArticle(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ArticleId))
	}
}
