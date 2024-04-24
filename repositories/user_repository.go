package repositories

//TODO: USER REPO

import (
	"sync"

	"goshort/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

var (
	user_repository_instance *UserRepository
	once_user_repository     sync.Once
)

func GetUserRepository() *UserRepository {
	once_user_repository.Do(func() {
		user_repository_instance = &UserRepository{
			collection: utils.GetMongoDb().Client.Database("goshort").Collection("users"),
		}
	})

	return user_repository_instance
}
