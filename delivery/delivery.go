package delivery

import (
	m "main/domain/model"
)

type ChatInterface interface {
	BroadcastMsg(msg *m.ClassBroadcastMessage) error
	SendNotification(msg *m.SingleMessage) error
}

type CalendarInterface interface {
	GetEvents(teacherID int) (m.CalendarEvents, error)
	CreateEvent(in m.CreateEvent) error
	// UpdateEvent(EventData) error
	DeleteEvent(calendarDB, eventID string) error
}
