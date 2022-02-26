package item

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpango/glg"

	"github.com/dannegm/spendzer-service/common/constants"
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
	glg.Info("Item > Controller")

	return func(c *gin.Context) {
		is, err := itemService.CreateItemService(config.Mongo.URI)

		if err != nil {
			fmt.Println(err)
		}

		name := strings.Trim(c.PostForm("name"), " ")

		if name == "" {
			c.JSON(constants.HTTP_BAD_REQUEST, ErrorResponse{
				Status:  constants.HTTP_BAD_REQUEST,
				Message: "El campo `name` no debe estar vacío",
			})
		}

		itemCreated, err := is.AddNewItem(itemDomain.Item{
			ID:   uuid.New(),
			Name: name,
		})

		if err != nil {
			c.JSON(constants.HTTP_INTERNAL_ERROR, ErrorResponse{
				Status:  constants.HTTP_INTERNAL_ERROR,
				Message: "Ah ocurrido un error desconocido",
			})
		}

		c.JSON(constants.HTTP_CREATED, ItemResponse{
			Data: itemCreated,
		})
	}
}
