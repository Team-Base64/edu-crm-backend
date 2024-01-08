package chat

import (
	"context"

	d "main/delivery"
	proto "main/delivery/grpc/chat/proto"
	m "main/domain/model"
)

type ChatService struct {
	client proto.ChatControllerClient
}

func NewChatService(c proto.ChatControllerClient) d.ChatInterface {
	return &ChatService{
		client: c,
	}
}

func (cs *ChatService) BroadcastMsg(msg *m.ClassBroadcastMessage) error {
	_, err := cs.client.BroadcastMsg(
		context.Background(),
		&proto.BroadcastMessage{
			ClassID:        int32(msg.ClassID),
			Title:          msg.Title,
			Description:    msg.Description,
			AttachmentURLs: msg.Attaches,
		})
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChatService) SendNotification(msg *m.SingleMessage) error {
	_, err := cs.client.SendNotification(
		context.Background(),
		&proto.Message{
			ChatID:         int32(msg.ChatID),
			Text:           msg.Text,
			AttachmentURLs: msg.Attaches,
		})
	if err != nil {
		return err
	}
	return nil
}
