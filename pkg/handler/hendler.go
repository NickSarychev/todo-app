package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) initRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}
	api := router.Group("/api")
	{
		lists := api.Group("/api")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group(":id/items")
			{
				items.POST("/")
				items.GET("/")
				items.GET("/:item_id")
				items.PUT("/:id")
				items.DELETE("/:id")
			}
		}
	}
	return router
}
