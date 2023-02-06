package usergin

import (
	"gin_form/common"
	userbusiness "gin_form/module/user/business"
	usermodel "gin_form/module/user/model"
	userstorage "gin_form/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter usermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Limit = 1

		paging.Validate()

		storage := userstorage.NewDbStore(db)
		business := userbusiness.ListUserBusiness(storage)

		result, err := business.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SuccessResponse(result, paging, filter))
	}
}
