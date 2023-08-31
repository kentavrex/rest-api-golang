package main

import (
	"golang-rest-api-template/controllers"
	"golang-rest-api-template/docs"
	"golang-rest-api-template/models"
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Rest api Golang Avito App API
// @version 1.0
// @description API Server for Segments Application

// @host localhost:8001
// @BasePath /

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	models.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", controllers.FindUsers)
		v1.POST("/users", controllers.CreateUser)
		v1.GET("/users/:id", controllers.FindUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)

		user_segments := v1.Group("/user_segments")
		{
			user_segments.GET("/", controllers.FindUserSegments)
			user_segments.POST("/", controllers.CreateUserSegment)
			user_segments.GET("/:id", controllers.FindUserSegment)
			user_segments.PUT("/:id", controllers.UpdateUserSegment)
			user_segments.DELETE("/:id", controllers.DeleteUserSegment)
		}

		v1.GET("/segments", controllers.FindSegments)
		v1.POST("/segments", controllers.CreateSegment)
		v1.GET("/segments/:id", controllers.FindSegment)
		v1.PUT("/segments/:id", controllers.UpdateSegment)
		v1.DELETE("/segments/:id", controllers.DeleteSegment)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
