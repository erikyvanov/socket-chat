package repositories

import (
	"context"
	"sync"
	"time"

	"github.com/erikyvanov/chat-fh/database"
	"github.com/erikyvanov/chat-fh/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessagesRepository struct {
	collection *mongo.Collection
}

var onceMessageRepository sync.Once
var messagesRepository *MessagesRepository

const messagesCollectionName = "messages"

func GetMessageRepository() *MessagesRepository {
	onceMessageRepository.Do(func() {
		messagesRepository = &MessagesRepository{
			collection: database.GetMongoDatabase().Collection(messagesCollectionName),
		}
	})

	return messagesRepository
}

func (mr *MessagesRepository) SaveMessage(message models.ChatMessage) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := mr.collection.InsertOne(ctx, message)
	return err
}

func (mr *MessagesRepository) GetChatMessages(user_email1, user_email2 string, page int) ([]models.ChatMessage, error) {
	var messages []models.ChatMessage
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	options := options.Find()
	options.SetSkip(int64(findLimit * page))
	options.SetLimit(findLimit)

	filter := bson.M{"$or": bson.A{bson.M{
		"sender_email":  user_email1,
		"reciver_email": user_email2,
	}, bson.M{
		"sender_email":  user_email2,
		"reciver_email": user_email1,
	}}}

	cursor, err := mr.collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var message models.ChatMessage
		if err = cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
