package link_repository

import (
	"context"
	"goshort/dtos/link_dto"
	"goshort/models"
	"goshort/utils/mongodb"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LinkRepository struct {
	collection *mongo.Collection
}

var (
	instance *LinkRepository
	once     sync.Once
)

func GetInstance() *LinkRepository {
	once.Do(func() {
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
	})
	return instance
}

// getters
func (lr *LinkRepository) GetCollection() *mongo.Collection {
	return lr.collection
}

//methods

func (lr *LinkRepository) AddLink(link *link_dto.LinkDTO_Info) (*link_dto.LinkDTO_Get, error) {

	user_id_primitive, err_casting := primitive.ObjectIDFromHex(link.UserID)

	if err_casting != nil {
		return nil, err_casting
	}

	newLink := models.Link{
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      user_id_primitive,
		Clicks:      0,
	}

	result, err := lr.collection.InsertOne(context.Background(), newLink)

	if err != nil {
		return nil, err
	}

	created_link, err := instance.GetLinkById(result.InsertedID.(primitive.ObjectID))

	if created_link == nil {
		return nil, err
	}

	return created_link, nil
}

func (lr *LinkRepository) updateLinkClicks(link_id primitive.ObjectID, clicks int) error {
	add_clicks := clicks + 1
	filter := bson.D{{Key: "_id", Value: link_id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "clicks", Value: add_clicks}}}}

	_, err := lr.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LinkRepository) DeleteLinkById(link_id primitive.ObjectID) (*link_dto.LinkDTO_Get, error) {
	var link models.Link

	result := lr.collection.FindOneAndDelete(context.Background(), bson.M{"_id": link_id})

	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&link); err != nil {
		return nil, err
	}

	returned_link := &link_dto.LinkDTO_Get{
		ID:          link.ID.Hex(),
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID.Hex(),
		Clicks:      link.Clicks,
	}

	return returned_link, nil
}

func (lr *LinkRepository) GetLinkById(id primitive.ObjectID) (*link_dto.LinkDTO_Get, error) {
	var link models.Link

	result := instance.collection.FindOne(context.Background(), bson.M{"_id": id})

	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&link); err != nil {
		return nil, err
	}

	link_returned := &link_dto.LinkDTO_Get{
		ID:          link.ID.Hex(),
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID.Hex(),
		Clicks:      link.Clicks,
	}

	return link_returned, nil
}

func (lr *LinkRepository) GetLinkByOriginalUrl(original_url string) (*link_dto.LinkDTO_Get, error) {
	var link models.Link

	result := instance.collection.FindOne(context.Background(), bson.M{"original_url": original_url})

	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&link); err != nil {
		return nil, err
	}

	link_returned := &link_dto.LinkDTO_Get{
		ID:          link.ID.Hex(),
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID.Hex(),
		Clicks:      link.Clicks,
	}

	return link_returned, nil
}

func (lr *LinkRepository) GetLinkByShortUrl(short_url string) (*link_dto.LinkDTO_Get, error) {
	var link models.Link

	result := instance.collection.FindOne(context.Background(), bson.M{"short_url": short_url})

	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(&link); err != nil {
		return nil, err
	}

	if err := instance.updateLinkClicks(link.ID, link.Clicks); err != nil {
		return nil, err
	}

	link_returned := &link_dto.LinkDTO_Get{
		ID:          link.ID.Hex(),
		ShortUrl:    link.ShortUrl,
		OriginalUrl: link.OriginalUrl,
		UserID:      link.UserID.Hex(),
		Clicks:      link.Clicks,
	}

	return link_returned, nil
}

func (lr *LinkRepository) GetLinksByUserId(user_id primitive.ObjectID) ([]*link_dto.LinkDTO_Get, error) {
	var links []*models.Link
	var links_dtos []*link_dto.LinkDTO_Get

	cursor, err := instance.collection.Find(context.Background(), bson.M{"user_id": user_id})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var link models.Link
		if err := cursor.Decode(&link); err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	for _, link := range links {
		link_dto := &link_dto.LinkDTO_Get{
			ID:          link.ID.Hex(),
			ShortUrl:    link.ShortUrl,
			OriginalUrl: link.OriginalUrl,
			UserID:      link.UserID.Hex(),
			Clicks:      link.Clicks,
		}

		links_dtos = append(links_dtos, link_dto)
	}

	return links_dtos, nil
}
