package http

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(hv ServerViewService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	view = hv

	// disable gin logger middleware
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(gin.Recovery())

	// set up routers
	r.POST("/signup", Signup)
	r.GET("/users/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		GetUser(c, userID)
	})
	r.PATCH("/users/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		PatchUser(c, userID)
	})
	r.POST("/close", Close)
	r.GET("/api/healthy", Healthy)
	return r
}
