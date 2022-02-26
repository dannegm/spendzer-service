package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dannegm/spendzer-service/config"
	"github.com/dannegm/spendzer-service/domain/item"
)

var MongoCollectionName string = "items"

type MongoRepository struct {
	db   *mongo.Database
	item *mongo.Collection
}

type mongoItem struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromItem(i item.Item) mongoItem {
	return mongoItem{
		ID:   i.GetID(),
		Name: i.GetName(),
	}
}

func (m mongoItem) ToAggregate() item.Item {
	i := item.Item{}

	i.SetID(m.ID)
	i.SetName(m.Name)

	return i
}

func NewRepository(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Mongo.Database)
	items := db.Collection(MongoCollectionName)

	return &MongoRepository{
		db:   db,
		item: items,
	}, nil
}

func (mr *MongoRepository) Add(i item.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.Mongo.Timeout*time.Second)
	defer cancel()

	internal := NewFromItem(i)
	_, err := mr.item.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}
