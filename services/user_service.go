package services

import (
	"errors"
	"sync"

	"github.com/erikyvanov/chat-fh/jwt"
	"github.com/erikyvanov/chat-fh/models"
	"github.com/erikyvanov/chat-fh/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *repositories.UserRepository
}

var once sync.Once
var userService *UserService

var (
	ErrEmailExists            = errors.New("the user is already registered")
	ErrUserDontExist          = errors.New("user does not exist")
	ErrPasswordOrEmailInvalid = errors.New("the email or password is not valid")
)

func GetUserService() *UserService {
	once.Do(func() {
		userService = &UserService{
			repository: repositories.GetUserRepository(),
		}
	})

	return userService
}

func (us *UserService) RegisterUser(user *models.User) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", err
	}

	jwt, err := jwt.GenerateJWT(*user)
	if err != nil {
		return "", err
	}

	user.Password = string(passwordHash)
	user.Online = false
	err = us.repository.Save(*user)
	if mongo.IsDuplicateKeyError(err) {
		return "", ErrEmailExists
	}
	user.Password = ""

	return jwt, err
}

func (us *UserService) Login(user *models.User) (string, error) {
	dbUser, err := us.repository.GetUser(user.Email)
	if err != nil {
		return "", ErrPasswordOrEmailInvalid
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return "", ErrPasswordOrEmailInvalid
	}

	jwt, err := jwt.GenerateJWT(*dbUser)
	if err != nil {
		return "", err
	}

	user.Name = dbUser.Name
	user.Online = dbUser.Online
	user.Password = ""

	return jwt, err
}

func (us *UserService) GetUser(email string) (*models.User, error) {
	user, err := us.repository.GetUser(email)
	if err != nil {
		return nil, ErrUserDontExist
	}

	user.Password = ""
	return user, nil
}
