package application

import (
	"MSS/src/application/iam"
	"MSS/src/application/message"
	"MSS/src/infrastructure/http"
	"MSS/src/infrastructure/mongo"
	infraiam "MSS/src/infrastructure/mongo/iam"
	inframessage "MSS/src/infrastructure/mongo/message"
)

type ServiceProvider struct {
	iamService     iam.Service
	messageService message.Service
}

func NewServiceProvider(opt *Options) *ServiceProvider {
	mongoClient := mongo.NewClient(opt.MongoUri)
	httpClient := http.NewClient()
	ur := infraiam.NewUserRepository(mongoClient)
	mr := inframessage.NewMessageRepository(mongoClient)
	ms := http.NewMessageService(httpClient)
	iamService := iam.NewService(ur)
	messageService := message.NewService(iamService, mr, ms)

	return &ServiceProvider{
		iamService:     iamService,
		messageService: messageService,
	}
}

func (s ServiceProvider) IamService() iam.Service {
	return s.iamService
}

func (s ServiceProvider) MessageService() message.Service {
	return s.messageService
}
