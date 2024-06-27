package main

import (
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
	client := mongo.NewClient("mongodb://localhost:27017")

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
	client := mongo.NewClient("mongodb://localhost:27017")

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
	client := mongo.NewClient("mongodb://localhost:27017")

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
