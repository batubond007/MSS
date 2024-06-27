package main

import (
	"MSS/src/application"
	"MSS/src/application/iam"
	"MSS/src/application/message"
	infrahttp "MSS/src/infrastructure/http"
	"MSS/src/infrastructure/mongo"
	infraiam "MSS/src/infrastructure/mongo/iam"
	inframessage "MSS/src/infrastructure/mongo/message"
)

func NewServiceProvider(opt *Options) *application.ServiceProvider {
	mongoClient := mongo.NewClient(opt.MongoUri)
	httpClient := infrahttp.NewClient()
	ur := infraiam.NewUserRepository(mongoClient)
	mr := inframessage.NewMessageRepository(mongoClient)
	ms := infrahttp.NewMessageService(httpClient)
	iamService := iam.NewService(ur)
	messageService := message.NewService(iamService, mr, ms)

	return application.NewServiceProvider(iamService, messageService)
}
