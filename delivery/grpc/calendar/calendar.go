package ctrl

import (
	"context"
	"time"

	d "main/delivery"
	proto "main/delivery/grpc/calendar/proto"
	e "main/domain/errors"
	m "main/domain/model"
)

type CalendarService struct {
	client proto.CalendarClient
}

func NewCalendarService(c proto.CalendarClient) d.CalendarInterface {
	return &CalendarService{
		client: c,
	}
}

func (cs *CalendarService) GetEvents(teacherID int) ([]m.CalendarEvent, error) {
	events, err := cs.client.GetEventsCalendar(
		context.Background(),
		&proto.GetEventsRequestCalendar{TeacherID: int32(teacherID)})
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	ans := []m.CalendarEvent{}
	for _, ev := range events.Events {
		t1, err := time.Parse(time.RFC3339, ev.StartDate)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		t2, err := time.Parse(time.RFC3339, ev.EndDate)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		ans = append(ans, m.CalendarEvent{
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

func (cs *CalendarService) CreateEvent(ev m.CalendarEvent, calID string) (string, error) {
	CreateEventResp, err := cs.client.CreateEvent(
		context.Background(), &proto.CreateEventRequest{CalendarID: calID,
			Event: &proto.EventData{Title: ev.Title,
				Description: ev.Description,
				StartDate:   ev.StartDate.Format(time.RFC3339Nano),
				EndDate:     ev.EndDate.Format(time.RFC3339Nano),
				Id:          ev.ID,
				ClassID:     int32(ev.ClassID)}})
	return CreateEventResp.EventID, err
}

func (cs *CalendarService) DeleteEvent(calendarDB, eventID string) error {
	_, err := cs.client.DeleteEvent(
		context.Background(), &proto.DeleteEventRequest{Id: eventID, CalendarID: calendarDB})
	return err
}

func (cs *CalendarService) UpdateEvent(ev m.CalendarEvent, calID string) error {
	_, err := cs.client.UpdateEvent(
		context.Background(), &proto.UpdateEventRequest{CalendarID: calID,
			Event: &proto.EventData{Title: ev.Title,
				Description: ev.Description,
				StartDate:   ev.StartDate.Format(time.RFC3339Nano),
				EndDate:     ev.EndDate.Format(time.RFC3339Nano),
				Id:          ev.ID,
				ClassID:     int32(ev.ClassID)}})
	return err
}

func (cs *CalendarService) CreateCalendar(teacherID int) error {
	_, err := cs.client.CreateCalendar(
		context.Background(), &proto.CreateCalendarRequest{TeacherID: int32(teacherID)})
	return err
}
