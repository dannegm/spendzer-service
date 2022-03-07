package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"

	"github.com/dannegm/spendzer-service/api/item"
	"github.com/dannegm/spendzer-service/api/user"
)

func RegisterRoutes(app *gin.Engine) {
	glg.Info("Mounting routes...")
	// Add your routes here

	item.Router(app)
	user.Router(app)
}
