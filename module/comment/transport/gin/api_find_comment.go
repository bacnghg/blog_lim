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

func GetComment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// id, err := strconv.Atoi(c.Param("user_id"))
		id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := commentstorage.NewDbStore(db)

		business := commentbusiness.NewFindCommentBusiness(storage)

		data, err := business.FindComment(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
