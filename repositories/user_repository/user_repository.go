package user_repository

import (
	"context"
	user_dtos "goshort/dtos/user_dto"
	"goshort/models"
	"goshort/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

var instance *UserRepository

func GetInstance() *UserRepository {
	if instance == nil {
		instance = &UserRepository{
			collection: mongodb.GetInstance().GetClient().Database("goshort").Collection("users"),
		}
	}
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

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	createdUser := instance.GetUserById(insertedID)

	if createdUser == nil {
		return nil, err
	}

	return createdUser, nil
}

func (ur *UserRepository) GetUserById(id primitive.ObjectID) *user_dtos.UserDTO_Info {
	var user models.User

	err := instance.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil
	}

	return &user_dtos.UserDTO_Info{
		Username: user.Username,
		Email:    user.Email,
		Created:  user.Created,
	}
}

func (ur *UserRepository) GetUserByEmail(email string) *user_dtos.UserDTO_Info {
	var user models.User

	err := instance.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil
	}

	return &user_dtos.UserDTO_Info{
		Username: user.Username,
		Email:    user.Email,
		Created:  user.Created,
	}
}
