package commentgin

import (
	"gin_form/common"
	commentbusiness "gin_form/module/comment/business"
	commentmodel "gin_form/module/comment/model"
	commentstorage "gin_form/module/comment/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCommentHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data commentmodel.CommentUpdate

		id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := commentstorage.NewDbStore(db)
		business := commentbusiness.UpdateCommentBusiness(storage)

		if err := business.UpdateCommentById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
