package chat

//пока не смог подключить

// import (
// 	context "context"
// )

// type ChatServiceInterface interface {
// 	BroadcastMsg(ctx context.Context, msg *BroadcastMessage) (Nothing, error)
// }

// // TODO Сюда идет grpc клиент, который получается после генерации proto файла
// // type ChatService struct{
// // 	client GrpcClientFromGeneration
// // }

// // func NewChatService(c *GrpcClientFromGeneration) ChatServiceInterface {
// // 	return &ChatService{
// // 		client: c,
// // 	}
// // }

// // func (cs *ChatService) BroadcastMsg(msg *model.ClassBroudcastMessage) error {
// // 	return nil
// // }

// type ChatService struct {
// 	client BotChatClient
// 	//ChatServiceInterface
// }

// func NewChatService(c BotChatClient) *ChatService {
// 	return &ChatService{
// 		client: c,
// 	}
// }

// func (cs *ChatService) BroadcastMsg(ctx context.Context, msg *BroadcastMessage) (Nothing, error) {
// 	_, err := cs.client.BroadcastMsg(
// 		context.Background(),
// 		&BroadcastMessage{
// 			ClassID:        int32(msg.ClassID),
// 			Title:          msg.Title,
// 			Description:    msg.Description + "\n + Дедлайн: " + msg.DeadlineTime.String(),
// 			AttachmentURLs: msg.Attaches,
// 		})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
