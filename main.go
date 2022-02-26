package main

import (
	"context"
	"fmt"

	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"github.com/dannegm/spendzer-service/config"
	"github.com/dannegm/spendzer-service/domain/item"
	"github.com/dannegm/spendzer-service/domain/item/mongo"
)

type ItemsService struct {
	Repository *mongo.MongoRepository
}

func createItemService(connectionString string) (*ItemsService, error) {
	repo, err := mongo.NewRepository(context.Background(), connectionString)

	if err != nil {
		return &ItemsService{}, err
	}

	return &ItemsService{
		Repository: repo,
	}, nil
}

func main() {
	fmt.Println("Happy hacking!!")

	is, err := createItemService(config.Mongo.URI)

	if err != nil {
		fmt.Println(err)
	}

	newItem, err := item.NewItem("Test 001")
	if err != nil {
		fmt.Println(err)
	}

	is.Repository.Add(newItem)
}

// cannot use value in struct literal
