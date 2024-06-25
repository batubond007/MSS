package iam

import (
	"MSS/src/domain/user"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const db = "default"
const collection = "user"

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) user.UserRepository {
	return UserRepository{client: client}
}

func (u UserRepository) FindByPhone(phone string) (user.User, error) {
	var res User
	row := u.collection().FindOne(context.Background(), bson.D{{"phone", phone}})
	err := row.Decode(&res)
	if err != nil {
		return user.User{}, err
	}
	return toDomainUser(res), nil
}

func (u UserRepository) Save(user user.User) error {
	_, err := u.collection().InsertOne(nil, bson.D{
		{"phone", user.Phone()},
		{"password", user.Password()}})
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Delete(user user.User) error {
	_, err := u.collection().DeleteOne(context.Background(), bson.D{{"phone", user.Phone()}})
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) List() ([]user.User, error) {
	var res []User
	rows, err := u.collection().Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for rows.Next(nil) {
		var row User
		err := rows.Decode(&row)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}
	return toDomainUserList(res), nil
}

func (u UserRepository) collection() *mongo.Collection {
	return u.client.Database(db).Collection(collection)
}
