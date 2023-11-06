package chat

//пока не смог подключить

import (
	context "context"
	"main/domain/model"
)

type ChatServiceInterface interface {
	BroadcastMsg(msg *model.ClassBroadcastMessage) error
}

type ChatService struct {
	client BotChatClient
}

func NewChatService(c BotChatClient) ChatServiceInterface {
	return &ChatService{
		client: c,
	}
}

func (cs *ChatService) BroadcastMsg(msg *model.ClassBroadcastMessage) error {
	_, err := cs.client.BroadcastMsg(
		context.Background(),
		&BroadcastMessage{
			ClassID:        int32(msg.ClassID),
			Title:          msg.Title,
			Description:    msg.Description, //+ "\n + Дедлайн: " + msg.DeadlineTime.String(),
			AttachmentURLs: msg.Attaches,
		})
	if err != nil {
		return err
	}
	return nil
}
