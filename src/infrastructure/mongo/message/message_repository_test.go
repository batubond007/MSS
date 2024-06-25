package message

import (
	"MSS/src/domain/message"
	"MSS/src/infrastructure/mongo"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func newRepo() *messageRepository {
	client := mongo.NewClient("mongodb://localhost:27017")
	return &messageRepository{client: client}
}

func mockMessage() Message {
	return Message{
		Id:      uuid.NewString(),
		Phone:   "test",
		Content: "test-content",
		Sent:    false,
	}
}

func TestMessageRepository_ListUnsentMessages(t *testing.T) {
	mr := newRepo()
	if _, err := mr.collection().InsertMany(context.Background(), []interface{}{
		mockMessage(),
		mockMessage(),
	}); err != nil {
		t.Errorf(err.Error())
		return
	}
	messages, err := mr.ListUnsentMessages("test", 2)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if _, err = mr.collection().DeleteMany(context.Background(), bson.D{{"phone", "test"}}); err != nil {
		t.Errorf(err.Error())
		return
	}
	if len(messages) != 2 {
		t.Error()
	}
}

func TestMessageRepository_UpdateMessage(t *testing.T) {
	mr := newRepo()
	msg := mockMessage()
	if _, err := mr.collection().InsertOne(context.Background(), msg); err != nil {
		t.Errorf(err.Error())
		return
	}
	dmsg := message.NewMessage(msg.Id, msg.Phone, msg.Content, true)
	if err := mr.UpdateMessage(dmsg); err != nil {
		t.Errorf(err.Error())
		return
	}
	var res Message
	if err := mr.collection().FindOne(context.Background(), bson.D{{"_id", msg.Id}}).Decode(&res); err != nil {
		t.Errorf(err.Error())
		return
	}
	if _, err := mr.collection().DeleteOne(context.Background(), bson.D{{"_id", msg.Id}}); err != nil {
		t.Errorf(err.Error())
		return
	}
	if res.Sent == msg.Sent {
		t.Error()
	}
}
