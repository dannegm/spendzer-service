package item

import (
	"context"

	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"github.com/dannegm/spendzer-service/domain/item"
	itemRepositories "github.com/dannegm/spendzer-service/domain/item/repositories"
)

type ItemsService struct {
	Repository *itemRepositories.MongoRepository
}

func CreateItemService(connectionString string) (*ItemsService, error) {
	repo, err := itemRepositories.NewRepository(context.Background(), connectionString)

	if err != nil {
		return &ItemsService{}, err
	}

	return &ItemsService{
		Repository: repo,
	}, nil
}

func (is *ItemsService) AddNewItem(newItem item.Item) (item.Item, error) {
	err := is.Repository.Add(newItem)
	return newItem, err
}
