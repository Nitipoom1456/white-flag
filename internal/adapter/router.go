package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/white-flag/internal/adapter/handler"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(
	router *gin.Engine,
) *Router {
	return &Router{
		router: router,
	}
}

func (r *Router) SetupRoutes(handler *handler.Handler) {
	r.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	v1 := r.router.Group("v1")

	apps := v1.Group("apps")
	apps.POST("", handler.CreateApp)
	apps.GET("", handler.ListApps)
	apps.POST("/:id/versions", handler.CreateAppVersion)
	apps.GET("/:id/versions", handler.GetAppVersionByAppID)
	apps.GET("/:id/environments", handler.GetEnvironmentByAppID)
	apps.POST("/:id/environments", handler.CreateEnvironment)

	appVersions := v1.Group("app-versions")
	appVersions.DELETE("/:id", handler.DeleteAppVersion)

	environments := v1.Group("environments")
	environments.DELETE("/:id", handler.DeleteEnvironment)
}
