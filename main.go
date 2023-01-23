package main

import (
	articlegin "gin_form/module/article/transport/gin"
	usergin "gin_form/module/user/transport/gin"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//dsn := "root:admin@123@tcp(localhost:3306)/cowell_tool?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root@tcp(localhost:3306)/bac_blog_v1?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL: ", err)
	}

	log.Println("Connected: ", db)

	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("/users/:user_id", usergin.GetUser(db))
		v1.PUT("/users/:user_id", usergin.UpdateUserHandler(db))
		v1.DELETE("/users/:user_id", usergin.DeleteUserHandler(db))
		v1.POST("/users", usergin.CreateUserHandler(db))
		v1.GET("/users", usergin.GetListUserHandler(db))

		v1.GET("/articles/:article_id", articlegin.GetArticle(db)) // get user with id
		v1.PUT("/articles/:article_id", articlegin.UpdateArticleHandler(db))
		v1.DELETE("/articles/:article_id", articlegin.DeleteArticleHandler(db))
		v1.POST("/articles", articlegin.CreateArticleHandler(db))
		v1.GET("/articles", articlegin.GetListArticleHandler(db)) // get all user

	}

	router.Run(":3000")
}
