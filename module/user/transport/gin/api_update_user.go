package usergin

import (
	"gin_form/common"
	userbusiness "gin_form/module/user/business"
	usermodel "gin_form/module/user/model"
	userstorage "gin_form/module/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserUpdate

		id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := userstorage.NewDbStore(db)
		business := userbusiness.UpdateUserBusiness(storage)

		if err := business.UpdateUserById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
