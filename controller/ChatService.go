package controller

import "main/domain/model"

type ChatServiceInterface interface {
	BroadcastMsg(msg *model.ClassBroudcastMessage) error
}

// TODO Сюда идет grpc клиент, который получается после генерации proto файла
// type ChatService struct{
// 	client GrpcClientFromGeneration
// }

// func NewChatService(c *GrpcClientFromGeneration) ChatServiceInterface {
// 	return &ChatService{
// 		client: c,
// 	}
// }

// func (cs *ChatService) BroadcastMsg(msg *model.ClassBroudcastMessage) error {
// 	return nil
// }
