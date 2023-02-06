package articlegin

import (
	"gin_form/common"
	articlebusiness "gin_form/module/article/business"
	articlemodel "gin_form/module/article/model"
	articlestorage "gin_form/module/article/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data articlemodel.ArticleUpdate

		id, err := strconv.Atoi(c.Param("article_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := articlestorage.NewDbStore(db)
		business := articlebusiness.UpdateArticleBusiness(storage)

		if err := business.UpdateArticleById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
