package usergin

import (
	"gin_form/common"
	userbusiness "gin_form/module/user/business"
	userstorage "gin_form/module/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := userstorage.NewDbStore(db)
		business := userbusiness.NewDeleteUser(storage)

		if err := business.DeleteUser(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
