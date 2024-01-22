package delivery

import (
	m "main/domain/model"
)

type ChatInterface interface {
	BroadcastMsg(msg *m.ClassBroadcastMessage) error
	SendNotification(msg *m.SingleMessage) error
}

type CalendarInterface interface {
	GetEvents(teacherID int) ([]m.CalendarEvent, error)
	CreateEvent(ev m.CalendarEvent, calendarDB string) (string, error)
	UpdateEvent(ev m.CalendarEvent, calendarDB string) error
	DeleteEvent(calendarDB, eventID string) error
	CreateCalendar(teacherID int) error
}
