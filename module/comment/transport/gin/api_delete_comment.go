package commentgin

import (
	"gin_form/common"
	commentbusiness "gin_form/module/comment/business"
	commentstorage "gin_form/module/comment/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteCommentHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := commentstorage.NewDbStore(db)
		business := commentbusiness.NewDeleteComment(storage)

		if err := business.DeleteComment(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
