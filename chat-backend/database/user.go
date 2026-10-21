package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string    `json:"username" bson:"username"`
	Password    string    `json:"-" bson:"password"`
	DisplayName string    `json:"displayName" bson:"displayName"`
	Avatar      string    `json:"avatar" bson:"avatar"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}

func InsertUser(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now().UTC()
	if user.ID == "" {
		user.ID = "user_" + bson.NewObjectID().Hex()
	}
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = now
	}

	_, err := GetCollection("go-chat", "users").InsertOne(ctx, user)
	return err
}

func FindUserByUsername(username string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := GetCollection("go-chat", "users").FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&user)
	return user, err
}

func EnsureUserIndexes(ctx context.Context) error {
	_, err := GetCollection("go-chat", "users").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}
