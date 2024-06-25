package main

import (
	"MSS/src/application"
	"MSS/src/infrastructure/mongo"
	"MSS/src/infrastructure/mongo/iam"
	"MSS/src/infrastructure/mongo/message"
	"context"
	"github.com/google/uuid"
	"os"
)

func mockUser(phone string) iam.User {
	return iam.User{
		Phone:    phone,
		Password: uuid.NewString(),
	}
}

func mockMessage(phone string) message.Message {
	return message.Message{
		Id:      uuid.NewString(),
		Phone:   phone,
		Content: "test-content",
		Sent:    false,
	}
}

func main() {
	switch os.Args[1] {
	case "Drop":
		Drop()
	case "InsertUser":
		InsertUser()
	case "InsertMessage":
		InsertMessage()
	}
}

func Drop() {
	opt := application.NewOptions()
	client := mongo.NewClient(opt.MongoUri)

	err := client.Database("default").Collection("user").Drop(context.Background())
	if err != nil {
		panic(err)
	}
	err = client.Database("default").Collection("message").Drop(context.Background())
	if err != nil {
		panic(err)
	}
}

func InsertUser() {
	opt := application.NewOptions()
	client := mongo.NewClient(opt.MongoUri)

	_, err := client.Database("default").Collection("user").InsertMany(context.Background(), []interface{}{
		mockUser("+905551111111"),
		mockUser("+905552222222"),
		mockUser("+905553333333"),
	})
	if err != nil {
		panic(err)
	}
}

func InsertMessage() {
	opt := application.NewOptions()
	client := mongo.NewClient(opt.MongoUri)

	_, err := client.Database("default").Collection("message").InsertMany(context.Background(), []interface{}{
		mockMessage("+905551111111"),
		mockMessage("+905551111111"),
		mockMessage("+905551111111"),
		mockMessage("+905551111111"),
		mockMessage("+905552222222"),
		mockMessage("+905552222222"),
		mockMessage("+905552222222"),
		mockMessage("+905552222222"),
		mockMessage("+905553333333"),
		mockMessage("+905553333333"),
		mockMessage("+905553333333"),
		mockMessage("+905553333333"),
	})
	if err != nil {
		panic(err)
	}
}
