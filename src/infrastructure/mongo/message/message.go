package message

import "MSS/src/domain/message"

type Message struct {
	Id      string `bson:"_id"`
	Phone   string `bson:"phone"`
	Content string `bson:"content"`
	Sent    bool   `bson:"sent"`
}

func toDomainMessage(m Message) message.Message {
	return message.NewMessage(m.Id, m.Phone, m.Content, m.Sent)
}
