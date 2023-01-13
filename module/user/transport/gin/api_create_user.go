package usergin

import (
	"gin_form/common"
	userbusiness "gin_form/module/user/business"
	usermodel "gin_form/module/user/model"
	userstorage "gin_form/module/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := userstorage.NewDbStore(db)
		business := userbusiness.NewCreateUser(storage)

		if err := business.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
