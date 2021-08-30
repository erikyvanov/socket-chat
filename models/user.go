package models

type User struct {
	Name     string `json:"name" validate:"required" bson:"name"`
	Email    string `json:"email" validate:"required,email" bson:"_id"`
	Password string `json:"password,omitempty" validate:"required,min=6" bson:"password"`
	Online   bool   `json:"online" bson:"online"`
}
