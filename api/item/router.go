package item

import (
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func Router(app *gin.Engine) {
	glg.Info("Item > Router")

	items := app.Group("/items")
	items.POST("/", AddNewItem())
}
