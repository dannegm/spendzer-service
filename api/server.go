package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

type ServerConfig struct {
	Port  int
	Debug bool
}

func Setup(s ServerConfig) {
	if !s.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	// Middlewares
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// Routes
	RegisterRoutes(app)

	// Config & Run
	hostname := fmt.Sprintf(":%d", s.Port)
	glg.Infof("Server running at %s", hostname)
	http.ListenAndServe(hostname, app)
	app.Run()
}
