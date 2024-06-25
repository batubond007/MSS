package http

import "MSS/src/domain/message"

type (
	Message struct {
		Phone   string `json:"phone"`
		Content string `json:"content"`
	}
	MessageListResponse struct {
		Messages []Message `json:"messages"`
	}
)

func NewMessageListResponse(messages []message.Message) MessageListResponse {
	resList := make([]Message, 0)
	for _, m := range messages {
		resList = append(resList, Message{
			Phone:   m.Phone(),
			Content: m.Content(),
		})
	}
	return MessageListResponse{
		Messages: resList,
	}
}
