package http

import (
	"MSS/src/domain/message"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const messageAddr = "https://webhook.site/b2beaace-f834-49c5-92c8-e6dc817444ac"
const authKey = "INS.me1x9uMcyYGlhKKQVPoc.bO3j9aZwRTOcA2Ywo"

type Message struct {
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

type messageService struct {
	client *http.Client
}

func NewMessageService(client *http.Client) message.MessageService {
	return &messageService{client: client}
}

func (m messageService) SendMessage(phone string, content string) error {
	data := Message{
		Phone:   phone,
		Content: content,
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("GET", messageAddr, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-ins-auth-key", authKey)
	response, err := m.client.Do(request)
	if err != nil {
		return err
	}
	if response.Status != "200 OK" {
		log.Printf("[Http Client] Sent message returned unexpected status: %s", response.Status)
	}
	return nil
}
