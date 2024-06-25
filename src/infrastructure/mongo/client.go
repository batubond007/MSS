package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(uri string) *mongo.Client {
	cred := options.Credential{Username: "root", Password: "password"}
	opt := options.Client().
		ApplyURI(uri).
		SetAuth(cred)
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	return client
}
