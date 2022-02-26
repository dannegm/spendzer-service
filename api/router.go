package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"

	"github.com/dannegm/spendzer-service/api/item"
)

func RegisterRoutes(app *gin.Engine) {
	glg.Info("Montando rutas")
	// Add your routes here

	item.MountRouter(app)
}
