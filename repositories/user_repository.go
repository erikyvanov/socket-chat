package repositories

import (
	"context"
	"sync"
	"time"

	"github.com/erikyvanov/chat-fh/database"
	"github.com/erikyvanov/chat-fh/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

const collectionName = "users"

var once sync.Once
var repository *UserRepository

func GetUserRepository() *UserRepository {
	once.Do(func() {
		repository = &UserRepository{
			collection: database.GetMongoDatabase().Collection(collectionName),
		}
	})

	return repository
}

func (ur *UserRepository) Save(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := ur.collection.InsertOne(ctx, user)
	return err
}

func (ur *UserRepository) GetUser(email string) (*models.User, error) {
	user := models.User{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := ur.collection.FindOne(ctx, bson.M{"_id": email}).Decode(&user)

	return &user, err
}
