package http

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(hv HayateViewService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	view = hv

	// disable gin logger middleware
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(gin.Recovery())

	// set up routers
	r.POST("/api/lookat", Lookat)
	r.POST("/api/heard", Heard)
	r.POST("/api/thought", Thought)
	r.POST("/api/heardVoice", HeardVoice)
	r.GET("/api/healthy", Healthy)
	return r
}
