package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

type MockResponse struct {
	Status  int    `json:"satus"`
	Message string `json:"message"`
}

func CreateNewUser() func(*gin.Context) {
	glg.Info("User > Controller > CreateNewUser")

	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, MockResponse{
			Status:  http.StatusCreated,
			Message: "Created",
		})
	}
}

func ReadAllUsers() func(*gin.Context) {
	glg.Info("User > Controller > ReadAllUsers")

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, MockResponse{
			Status:  http.StatusOK,
			Message: "Readed",
		})
	}
}

func ReadUserById() func(*gin.Context) {
	glg.Info("User > Controller > ReadUserById")

	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, MockResponse{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("Readed with ID: %s", id),
		})
	}
}

func UpdateUserById() func(*gin.Context) {
	glg.Info("User > Controller > UpdateUserById")

	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, MockResponse{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("Updated with ID: %s", id),
		})
	}
}

func RemoveUserById() func(*gin.Context) {
	glg.Info("User > Controller > RemoveUserById")

	return func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	}
}
