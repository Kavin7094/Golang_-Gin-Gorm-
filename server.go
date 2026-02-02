package main

import (
	"Learn_GO/controllers"
	"Learn_GO/initilaizers"

	"github.com/gin-gonic/gin"
)

func init() {
	initilaizers.LoadEnvVariable()
	initilaizers.ConnnectionToDB()
}

func main() {
	Server := gin.Default()

	Server.POST("sample-test/", controllers.PostsCreate)
	Server.GET("sample-test/", controllers.PostsIndex)
	Server.GET("sample-test/:id", controllers.PostScope)
	Server.PUT("sample-test/:id", controllers.PostUpdate)
	Server.DELETE("sample-test/:id", controllers.PostsDelete)

	Server.Run()
}
