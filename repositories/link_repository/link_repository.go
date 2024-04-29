package link_repository

import (
	"context"
	"goshort/dtos/link_dto"
	"goshort/models"
	"goshort/utils/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LinkRepository struct {
	collection *mongo.Collection
}

var instance *LinkRepository

func GetInstance() *LinkRepository {
	if instance == nil {
		new_link_repo := &LinkRepository{
			collection: mongodb.GetInstance().GetClient().Database("goshort").Collection("links"),
		}

		keys := bson.D{{Key: "short_url", Value: 1}, {Key: "original_url", Value: 1}, {Key: "user_id", Value: 1}}

		uniqueFields := mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetUnique(true),
		}

		_, err := new_link_repo.collection.Indexes().CreateOne(context.Background(), uniqueFields)

		if err != nil {
			panic(err)
		}

		instance = new_link_repo
	}
	return instance
}

// getters
func (lr *LinkRepository) GetCollection() *mongo.Collection {
	return lr.collection
}

//methods

func (lr *LinkRepository) AddLink(link *link_dto.LinkDTO_Info) (*link_dto.LinkDTO_Get, error) {
	newLink := models.Link{
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID,
		Clicks:      0,
	}

	result, err := lr.collection.InsertOne(context.Background(), newLink)

	created_link := instance.GetLinkById(result.InsertedID.(primitive.ObjectID))

	if created_link == nil {
		return nil, err
	}

	return created_link, nil
}

func (lr *LinkRepository) GetLinkById(id primitive.ObjectID) *link_dto.LinkDTO_Get {
	var link models.Link

	err := instance.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&link)
	if err != nil {
		return nil
	}

	return &link_dto.LinkDTO_Get{
		ID:          link.ID.Hex(),
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID,
		Clicks:      link.Clicks,
	}
}

func (lr *LinkRepository) GetLinkByOriginalUrl(original_url string) *link_dto.LinkDTO_Get {
	var link models.Link

	err := instance.collection.FindOne(context.Background(), bson.M{"original_url": original_url}).Decode(&link)

	if err != nil {
		return nil
	}

	return &link_dto.LinkDTO_Get{
		ID:          link.ID.Hex(),
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID,
		Clicks:      link.Clicks,
	}
}

func (lr *LinkRepository) GetLinkByShortUrl(short_url string) *link_dto.LinkDTO_Get {
	var link models.Link

	err := instance.collection.FindOne(context.Background(), bson.M{"short_url": short_url}).Decode(&link)

	if err != nil {
		return nil
	}

	return &link_dto.LinkDTO_Get{
		ID:          link.ID.Hex(),
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID,
		Clicks:      link.Clicks,
	}
}
