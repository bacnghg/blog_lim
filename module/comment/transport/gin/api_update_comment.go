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

		id, err := strconv.Atoi(c.Param("comment_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(http.StatusBadRequest,err.Error()))
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(http.StatusBadRequest,err.Error()))
			return
		}

		storage := commentstorage.NewDbStore(db)
		business := commentbusiness.UpdateCommentBusiness(storage)

		if err := business.UpdateCommentBusiness(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(http.StatusBadRequest,err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
