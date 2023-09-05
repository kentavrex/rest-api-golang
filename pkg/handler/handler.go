package handler

import (
	_ "avito-dynamic-segment-back/docs"
	"avito-dynamic-segment-back/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.POST("/", h.createUser)
		users.GET("/", h.getAllUsers)
		users.GET("/:id", h.getUserById)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)

		segments := users.Group(":id/segments")
		{
			segments.POST("/", h.addUserSegments)
			segments.GET("/", h.getUserSegments)
			//segments.DELETE("/", h.deleteUserSegments)
		}
	}

	segments := router.Group("/segments")
	{
		segments.POST("/", h.createSegment)
		segments.GET("/", h.getAllSegments)
		segments.GET("/:id", h.getSegmentById)
		segments.PUT("/:id", h.updateSegment)
		segments.DELETE("/:id", h.deleteSegment)
	}

	return router
}
