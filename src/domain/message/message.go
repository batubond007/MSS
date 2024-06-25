package message

type (
	MessageService interface {
		SendMessage(phone string, content string) error
	}
	MessageRepository interface {
		ListSentMessages(phone string, size int) ([]Message, error)
		ListUnsentMessages(phone string, size int) ([]Message, error)
		UpdateMessage(message Message) error
	}
	Message struct {
		id      string
		phone   string
		content string
		sent    bool
	}
)

func NewMessage(id string, phone string, content string, sent bool) Message {
	return Message{id: id, phone: phone, content: content, sent: sent}
}

func (m Message) Id() string {
	return m.id
}

func (m Message) Phone() string {
	return m.phone
}

func (m Message) Content() string {
	return m.content
}

func (m Message) Sent() bool {
	return m.sent
}
