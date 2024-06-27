package application

import (
	"MSS/src/application/iam"
	"MSS/src/application/message"
)

type ServiceProvider struct {
	iamService     iam.Service
	messageService message.Service
}

func NewServiceProvider(iamService iam.Service, messageService message.Service) *ServiceProvider {
	return &ServiceProvider{iamService: iamService, messageService: messageService}
}

func (s ServiceProvider) IamService() iam.Service {
	return s.iamService
}

func (s ServiceProvider) MessageService() message.Service {
	return s.messageService
}
