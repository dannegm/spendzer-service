package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func Router(app *gin.Engine) {
	glg.Info("User > Router")

	users := app.Group("/users")
	users.POST("/", CreateNewUser())
	users.GET("/", ReadAllUsers())
	users.GET("/:id", ReadUserById())
	users.PUT("/:id", UpdateUserById())
	users.DELETE("/:id", RemoveUserById())
}
