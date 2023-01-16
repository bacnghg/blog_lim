package commentgin

import (
	"gin_form/common"
	commentbusiness "gin_form/module/comment/business"
	commentmodel "gin_form/module/comment/model"
	commentstorage "gin_form/module/comment/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCommentHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data commentmodel.CommentCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := commentstorage.NewDbStore(db)
		business := commentbusiness.NewCreateComment(storage)

		if err := business.CreateComment(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.CommentId))
	}
}
