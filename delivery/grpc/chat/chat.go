package chat

import (
	"context"

	d "main/delivery"
	proto "main/delivery/grpc/chat/proto"
	e "main/domain/errors"
	m "main/domain/model"
)

type ChatService struct {
	client proto.ChatClient
}

func NewChatService(c proto.ChatClient) d.ChatInterface {
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
		return e.StacktraceError(err)
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
		return e.StacktraceError(err)
	}
	return nil
}
