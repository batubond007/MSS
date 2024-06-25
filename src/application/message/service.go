package message

import (
	"MSS/src/application/iam"
	"MSS/src/domain/message"
	domainiam "MSS/src/domain/user"
	"log"
	"time"
)

const sentMessageCount = 2
const numOfMessageRetrievalRetries = 2

type (
	Service interface {
		Start()
		ListSentMessages(user domainiam.User, page int) ([]message.Message, error)
	}
	service struct {
		iamService iam.Service
		mr         message.MessageRepository
		ms         message.MessageService
	}
)

func NewService(iamService iam.Service, mr message.MessageRepository, ms message.MessageService) Service {
	return &service{iamService: iamService, mr: mr, ms: ms}
}

func (s service) ListSentMessages(user domainiam.User, page int) ([]message.Message, error) {
	messages, err := s.mr.ListSentMessages(user.Phone(), page)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (s service) Start() {
	go func() {
		for {
			s.ProcessAllUsers()
			time.Sleep(10 * time.Second)
		}
	}()
}

func (s service) ProcessAllUsers() {
	list, err := s.iamService.List()
	if err != nil {
		log.Printf("[Message Service] Error while getting user list: %s", err.Error())
		return
	}
	for _, user := range list {
		if err = s.ProcessUser(user); err != nil {
			log.Printf("[Message Service] Unable to process user %s : %s", user.Phone(), err.Error())
		}
	}
}

func (s service) ProcessUser(user domainiam.User) error {
	messages, err := s.mr.ListUnsentMessages(user.Phone(), sentMessageCount)
	if err != nil && messages == nil {
		return err
	}
	// Retry certain times to fill the expected message count
	i := 0
	for err != nil && len(messages) < sentMessageCount && i < numOfMessageRetrievalRetries {
		leftMessages := sentMessageCount - len(messages)
		var newMessages []message.Message
		newMessages, err = s.mr.ListUnsentMessages(user.Phone(), leftMessages*2)
		messages = append(messages, newMessages...)
		i++
	}

	if len(messages) >= sentMessageCount {
		messages = messages[:sentMessageCount]
	}

	for _, msg := range messages {
		if err := s.ms.SendMessage(user.Phone(), msg.Content()); err != nil {
			log.Printf("[Message Service] Unable to send message %s : %s", msg.Id(), err.Error())
			continue
		}
		if err := s.mr.UpdateMessage(message.NewMessage(msg.Id(), msg.Phone(), msg.Content(), true)); err != nil {
			log.Printf("[Message Service] Unable to update message %d : %s", msg.Id(), err.Error())
			continue
		}
	}
	return nil
}
