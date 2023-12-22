package ctrl

//пока не смог подключить

import (
	context "context"
	"log"
	"main/domain/model"
	"time"
)

type CtrlServiceInterface interface {
	BroadcastMsg(msg *model.ClassBroadcastMessage) error
	SendNotification(msg *model.SingleMessage) error
	GetEvents(teacherID int) (model.CalendarEvents, error)
	// CreateEvent(EventData) error
	// UpdateEvent(EventData) error
	// DeleteEvent(DeleteEventRequest) error
}

type CtrlService struct {
	chatClient     ChatControllerClient
	calendarClient CalendarControllerClient
}

func NewCtrlService(c1 ChatControllerClient, c2 CalendarControllerClient) CtrlServiceInterface {
	return &CtrlService{
		chatClient:     c1,
		calendarClient: c2,
	}
}

func (cs *CtrlService) BroadcastMsg(msg *model.ClassBroadcastMessage) error {
	_, err := cs.chatClient.BroadcastMsg(
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

func (cs *CtrlService) SendNotification(msg *model.SingleMessage) error {
	_, err := cs.chatClient.SendNotification(
		context.Background(),
		&Message{
			ChatID:         int32(msg.ChatID),
			Text:           msg.Text,
			AttachmentURLs: msg.Attaches,
		})
	if err != nil {
		return err
	}
	return nil
}

func (cs *CtrlService) GetEvents(teacherID int) (model.CalendarEvents, error) {
	events, err := cs.calendarClient.GetEvents(
		context.Background(),
		&GetEventsRequest{TeacherID: int32(teacherID)})
	if err != nil {
		return model.CalendarEvents{}, err
	}
	ans := model.CalendarEvents{}
	for _, ev := range events.Events {
		t1, err := time.Parse(time.RFC3339, ev.StartDate)
		if err != nil {
			log.Println("err converting time")
			return model.CalendarEvents{}, err
		}
		t2, err := time.Parse(time.RFC3339, ev.EndDate)
		if err != nil {
			log.Println("err converting time")
			return model.CalendarEvents{}, err
		}
		ans.Events = append(ans.Events, model.CalendarEvent{ID: ev.Id, Title: ev.Title, Description: ev.Description, StartDate: t1, EndDate: t2, ClassID: int(ev.ClassID)})
	}
	return ans, nil
}
