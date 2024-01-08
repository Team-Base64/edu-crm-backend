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
