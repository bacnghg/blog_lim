package articlegin

import (
	"gin_form/common"
	articlebusiness "gin_form/module/article/business"
	articlestorage "gin_form/module/article/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := articlestorage.NewDbStore(db)
		business := articlebusiness.NewDeleteArticle(storage)

		if err := business.DeleteArticle(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
