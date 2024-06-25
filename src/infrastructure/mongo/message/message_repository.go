package message

import (
	"MSS/src/domain/message"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const db = "default"
const collection = "message"

type messageRepository struct {
	client *mongo.Client
}

func NewMessageRepository(client *mongo.Client) message.MessageRepository {
	return &messageRepository{client: client}
}

func (m messageRepository) ListSentMessages(phone string, size int) ([]message.Message, error) {
	var messages []message.Message
	cursor, err := m.collection().Find(
		context.Background(),
		bson.M{"phone": phone, "sent": true},
		options.Find().SetLimit(int64(size)),
	)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var msg Message
		if err = cursor.Decode(&msg); err != nil {
			log.Printf("[Message Repository] Unable to decode message for user : %s", phone)
			continue
		}
		messages = append(messages, toDomainMessage(msg))
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	if err = cursor.Close(context.Background()); err != nil {
		return nil, err
	}
	return messages, nil
}

func (m messageRepository) ListUnsentMessages(phone string, size int) ([]message.Message, error) {
	var messages []message.Message
	cursor, err := m.collection().Find(
		context.Background(),
		bson.M{"phone": phone, "sent": false},
		options.Find().SetLimit(int64(size)),
	)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var msg Message
		// If error occurs while decoding continue to next message to prevent blocking the message service
		if err = cursor.Decode(&msg); err != nil {
			log.Printf("[Message Repository] Unable to decode message for user : %s", phone)
			continue
		}
		messages = append(messages, toDomainMessage(msg))
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	if err = cursor.Close(context.Background()); err != nil {
		return nil, err
	}
	return messages, nil
}

func (m messageRepository) UpdateMessage(message message.Message) error {
	filter := bson.M{"_id": message.Id()}
	update := bson.M{"$set": bson.M{
		"phone":   message.Phone(),
		"content": message.Content(),
		"sent":    message.Sent(),
	}}
	_, err := m.collection().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m messageRepository) collection() *mongo.Collection {
	return m.client.Database(db).Collection(collection)
}
