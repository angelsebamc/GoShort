package user_repository

import (
	"context"
	user_dtos "goshort/dtos/user_dto"
	"goshort/models"
	"goshort/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
}

var instance *UserRepository

func GetInstance() *UserRepository {
	if instance == nil {
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

	createdUser := instance.GetUserById(result.InsertedID.(primitive.ObjectID))

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
		ID:       user.ID.Hex(),
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
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Created:  user.Created,
	}
}
