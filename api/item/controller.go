package item

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpango/glg"

	"github.com/dannegm/spendzer-service/config"
	itemDomain "github.com/dannegm/spendzer-service/domain/item"
	itemService "github.com/dannegm/spendzer-service/services/item"
)

type ItemResponse struct {
	Data itemDomain.Item `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"satus"`
	Message string `json:"message"`
}

func AddNewItem() func(*gin.Context) {
	glg.Info("Item > Controller > AddNewItem")

	return func(c *gin.Context) {
		is, err := itemService.CreateItemService(config.Mongo.URI)

		if err != nil {
			fmt.Println(err)
		}

		name := strings.Trim(c.PostForm("name"), " ")

		if name == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "El campo `name` no debe estar vacío",
			})
		}

		itemCreated, err := is.AddNewItem(itemDomain.Item{
			ID:   uuid.New(),
			Name: name,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Ah ocurrido un error desconocido",
			})
		}

		c.JSON(http.StatusCreated, ItemResponse{
			Data: itemCreated,
		})
	}
}
