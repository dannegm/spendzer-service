package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dannegm/spendzer-service/config"
	"github.com/dannegm/spendzer-service/domain/user"
)

var MongoCollectionName string = "users"

type MongoRepository struct {
	db    *mongo.Database
	users *mongo.Collection
}

type mongoUser struct {
	ID       uuid.UUID `bson:"id"`
	Username string    `bson:"username"`
	Password string    `bson:"password"`
	JWT      string    `bson:"jwt"`
}

func MapUserToMongo(u user.User) mongoUser {
	return mongoUser{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		JWT:      u.JWT,
	}
}

func (m mongoUser) ToUser() user.User {
	u := user.User{
		ID:       m.ID,
		Username: m.Username,
		Password: m.Password,
		JWT:      m.JWT,
	}

	return u
}

func NewRepository(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Mongo.Database)
	users := db.Collection(MongoCollectionName)

	return &MongoRepository{
		db:    db,
		users: users,
	}, nil
}

func (mr *MongoRepository) Create(u user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.Mongo.Timeout*time.Second)
	defer cancel()

	internal := MapUserToMongo(u)
	_, err := mr.users.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}
