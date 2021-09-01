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

const findLimit = 10

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

func (ur *UserRepository) SetUserConectionStatus(email string, status bool) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{"_id": email}
	update := bson.M{"$set": bson.M{"online": status}}

	_, err := ur.collection.UpdateOne(ctx, filter, update)

	return err
}

func (ur *UserRepository) GetAllUsersExceptUserRequest(userRequestEmail string, page int) ([]*models.User, error) {
	var users []*models.User

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	options := options.Find()
	options.SetSkip(int64(findLimit * page))
	options.SetLimit(findLimit)
	options.SetSort(bson.M{"online": -1})

	filter := bson.M{"_id": bson.M{"$ne": userRequestEmail}}

	cursor, err := ur.collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user models.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}

		user.Password = ""
		users = append(users, &user)
	}

	return users, nil
}
