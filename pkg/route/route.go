package route

import (
	"Todo/models"
	"Todo/pkg/action"
	"Todo/server"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Route struct {
	action *action.Action
}

func NewRoute(action *action.Action) *Route {
	return &Route{action: action}
}

func (r *Route) InitHTTPRoutes(config *models.ServerConfig) *gin.Engine {
	ginSetMode(config.ServerMode)
	router := gin.Default()
	allowOrigins := strings.Split(config.Domain, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"X-Session-ID", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/todo", r.Create)
		api.GET("/todo", r.Get)
		api.DELETE("/todo/:id", r.Delete)
		api.PUT("/todo/:id", r.Update)
		api.PUT("/todo/:id/done", r.UpdateDone)
	}

	return router
}

func ginSetMode(serverMode string) {
	if serverMode == server.DEVELOPMENT {
		gin.SetMode(gin.ReleaseMode)
	}
}
