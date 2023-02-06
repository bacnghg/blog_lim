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

func GetListCommentHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(http.StatusBadRequest,err.Error()))
			return
		}

		var filter commentmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(http.StatusBadRequest,err.Error()))
			return
		}

		paging.Validate()

		storage := commentstorage.NewDbStore(db)
		business := commentbusiness.ListCommentBusiness(storage)

		result, err := business.ListComment(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.SimpleErrorResponse(http.StatusBadRequest,err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.SuccessResponse(result, paging, filter))
	}
}
