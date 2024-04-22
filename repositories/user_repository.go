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
	once                     sync.Once
)

func GetUserRepository() *UserRepository {
	once.Do(func() {
		user_repository_instance = &UserRepository{
			collection: utils.GetMongoDbClient().GetClient().Database("mydatabase").Collection("users"),
		}
	})

	return user_repository_instance
}
