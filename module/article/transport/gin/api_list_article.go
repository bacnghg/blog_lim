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

func GetListArticleHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter articlemodel.FilterArticle

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Validate()

		storage := articlestorage.NewDbStore(db)
		business := articlebusiness.ListArticleBusiness(storage)

		result, err := business.ListArticle(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SuccessResponse(result, paging, filter))
	}
}
