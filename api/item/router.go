package item

import (
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func MountRouter(app *gin.Engine) {
	glg.Info("Item > Router")

	items := app.Group("/items")
	items.POST("/", AddNewItem())
}
