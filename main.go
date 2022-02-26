package main

import (
	"fmt"

	"github.com/google/uuid"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"github.com/dannegm/spendzer-service/config"
	item "github.com/dannegm/spendzer-service/domain/item"
	itemService "github.com/dannegm/spendzer-service/services/item"
)

func main() {
	is, err := itemService.CreateItemService(config.Mongo.URI)

	if err != nil {
		fmt.Println(err)
	}

	newItem := item.Item{
		ID:   uuid.New(),
		Name: "Holi",
	}
	is.AddNewItem(newItem)
}
