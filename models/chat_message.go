package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatMessage struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	SenderEmail  string             `json:"sender_email" bson:"sender_email"`
	ReciverEmail string             `json:"reciver_email" bson:"reciver_email"`
	Message      string             `json:"message" bson:"message"`
}
