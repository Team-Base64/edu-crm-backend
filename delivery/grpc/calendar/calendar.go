package ctrl

import (
	"context"
	"log"
	"time"

	d "main/delivery"
	proto "main/delivery/grpc/calendar/proto"
	m "main/domain/model"
)

type CalendarService struct {
	client proto.CalendarControllerClient
}

func NewCalendarService(c proto.CalendarControllerClient) d.CalendarInterface {
	return &CalendarService{
		client: c,
	}
}

func (cs *CalendarService) GetEvents(teacherID int) (m.CalendarEvents, error) {
	events, err := cs.client.GetEventsCalendar(
		context.Background(),
		&proto.GetEventsRequestCalendar{TeacherID: int32(teacherID)})
	if err != nil {
		return m.CalendarEvents{}, err
	}
	ans := m.CalendarEvents{}
	for _, ev := range events.Events {
		t1, err := time.Parse(time.RFC3339, ev.StartDate)
		if err != nil {
			log.Println("err converting time")
			return m.CalendarEvents{}, err
		}
		t2, err := time.Parse(time.RFC3339, ev.EndDate)
		if err != nil {
			log.Println("err converting time")
			return m.CalendarEvents{}, err
		}
		ans.Events = append(ans.Events, m.CalendarEvent{
			ID:          ev.Id,
			Title:       ev.Title,
			Description: ev.Description,
			StartDate:   t1,
			EndDate:     t2,
			ClassID:     int(ev.ClassID),
		})
	}
	return ans, nil
}

func (cs *CalendarService) CreateEvent(req m.CreateEvent) error {
	_, err := cs.client.CreateEvent(
		context.Background(), &proto.CreateEventRequest{CalendarID: req.CalendarID,
			Event: &proto.EventData{Title: req.Event.Title,
				Description: req.Event.Description,
				StartDate:   req.Event.StartDate.Format(time.RFC3339Nano),
				EndDate:     req.Event.EndDate.Format(time.RFC3339Nano),
				Id:          req.Event.ID,
				ClassID:     int32(req.Event.ClassID)}})
	return err
}

func (cs *CalendarService) DeleteEvent(calendarDB, eventID string) error {
	_, err := cs.client.DeleteEvent(
		context.Background(), &proto.DeleteEventRequest{Id: eventID, CalendarID: calendarDB})
	return err
}
