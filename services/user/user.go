package User

import (
	"context"

	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	user "github.com/dannegm/spendzer-service/domain/user"
	userRepositories "github.com/dannegm/spendzer-service/domain/user/repositories"
)

type UsersService struct {
	Repository *userRepositories.MongoRepository
}

func CreateUserService(connectionString string) (*UsersService, error) {
	repo, err := userRepositories.NewRepository(context.Background(), connectionString)

	if err != nil {
		return &UsersService{}, err
	}

	return &UsersService{
		Repository: repo,
	}, nil
}

func (us *UsersService) CreateNewUser(newUser user.User) (user.User, error) {
	err := us.Repository.Create(newUser)
	return newUser, err
}
