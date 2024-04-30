package user_repository

import (
	"context"
	user_dtos "goshort/dtos/user_dto"
	"goshort/models"
	"goshort/utils/mongodb"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

var (
	instance *UserRepository
	once     sync.Once
)

func GetInstance() *UserRepository {
	once.Do(func() {
		new_user_repo := &UserRepository{
			collection: mongodb.GetInstance().GetClient().Database("goshort").Collection("users"),
		}

		uniqueEmail := mongo.IndexModel{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		}

		_, err := new_user_repo.collection.Indexes().CreateOne(context.Background(), uniqueEmail)

		if err != nil {
			panic(err)
		}

		instance = new_user_repo
	})
	return instance
}

// getters
func (ur *UserRepository) GetCollection() *mongo.Collection {
	return ur.collection
}

// methods
func (ur *UserRepository) CreateUser(user *models.User) (*user_dtos.UserDTO_Info, error) {
	result, err := instance.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	createdUser, err := instance.GetUserById(result.InsertedID.(primitive.ObjectID))

	if err != nil {
		return nil, err
	}

	if createdUser == nil {
		return nil, err
	}

	return createdUser, nil
}

// TODO: Do a generic method to retreive data by attributes
func (ur *UserRepository) GetUserById(id primitive.ObjectID) (*user_dtos.UserDTO_Info, error) {
	var user models.User

	result := instance.collection.FindOne(context.Background(), bson.M{"_id": id})

	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	returned_user := &user_dtos.UserDTO_Info{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Created:  user.Created,
	}

	return returned_user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*user_dtos.UserDTO_Info, error) {
	var user models.User

	result := instance.collection.FindOne(context.Background(), bson.M{"email": email})

	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	returned_user := &user_dtos.UserDTO_Info{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Created:  user.Created,
	}

	return returned_user, nil
}
