package backend

import (
	"errors"
	"fmt"
	"log"
	"strings"

	e "main/domain/errors"
	m "main/domain/model"
	u "main/domain/utils"
)

func (uc *BackendUsecase) GetCalendarParams(teacherID int) (*m.CalendarParams, error) {
	params, err := uc.dataStore.GetCalendar(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return params, nil
}

func (uc *BackendUsecase) CreateCalendar(teacherID int) error {
	if err := uc.calendar.CreateCalendar(teacherID); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (uc *BackendUsecase) CreateCalendarEvent(req *m.CalendarEvent, teacherID int) error {
	calendarDB, err := uc.dataStore.GetCalendar(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}
	eventID, err := uc.calendar.CreateEvent(*req, calendarDB.InternalApiID)
	if err != nil {
		log.Println("Unable to create event: ", err)
		return e.StacktraceError(err)
	}

	bcMsg := m.ClassBroadcastMessage{
		ClassID: req.ClassID,
		Title:   "Новое событие!" + "\n" + req.Title,
		Description: req.Description + "\n" +
			"Начало: " + u.TimeToString(req.StartDate) + "\n" +
			"Окончание: " + u.TimeToString(req.EndDate) + "\n" +
			"Ссылка на календарь: " +
			"https://calendar.google.com/calendar/embed?ctz=Europe%2FMoscow&hl=ru&src=" + calendarDB.InternalApiID,
		Attaches: []string{},
	}
	if err := uc.chat.BroadcastMsg(&bcMsg); err != nil {
		err = uc.calendar.DeleteEvent(calendarDB.InternalApiID, eventID)
		if err != nil {
			log.Println("Unable to delete event after bc error: ", err)
		}
		log.Println("Unable to broadcast msg: ", err)
		return e.StacktraceError(err)
	}
	return nil
}

func (uc *BackendUsecase) GetCalendarEvents(teacherID int) ([]m.CalendarEvent, error) {
	events, err := uc.calendar.GetEvents(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return events, nil
}

func (uc *BackendUsecase) DeleteCalendarEvent(teacherID int, eventID string) error {
	calendarDB, err := uc.dataStore.GetCalendar(teacherID)
	if err != nil {
		return e.StacktraceError(err)
	}
	return uc.calendar.DeleteEvent(calendarDB.InternalApiID, eventID)
}

func (uc *BackendUsecase) UpdateCalendarEvent(req *m.CalendarEvent, teacherID int) error {
	s := strings.Split(req.Title, " ")
	newTitle := ""
	if len(s) > 2 && s[len(s)-2] == "Class" {
		newTitle = strings.Join(s[:len(s)-2], " ")
	} else {
		newTitle = req.Title
	}
	req.Title = newTitle + " Class " + fmt.Sprintf("%d", req.ClassID)

	calendarDB, err := uc.dataStore.GetCalendar(teacherID)
	if err != nil {
		return e.StacktraceError(err)
	}
	err = uc.calendar.UpdateEvent(*req, calendarDB.InternalApiID)
	if err != nil {
		return e.StacktraceError(err)
	}

	bcMsg := m.ClassBroadcastMessage{
		ClassID: req.ClassID,
		Title:   "Событие обновлено!" + "\n" + req.Title,
		Description: req.Description + "\n" +
			"Начало: " + u.TimeToString(req.StartDate) + "\n" +
			"Окончание: " + u.TimeToString(req.EndDate) + "\n" +
			"Ссылка на календарь: " +
			"https://calendar.google.com/calendar/embed?ctz=Europe%2FMoscow&hl=ru&src=" + calendarDB.InternalApiID,
		Attaches: []string{},
	}
	if err := uc.chat.BroadcastMsg(&bcMsg); err != nil {
		err = uc.calendar.DeleteEvent(calendarDB.InternalApiID, req.ID)
		if err != nil {
			err = e.StacktraceError(err, errors.New("unable to delete event after bc error"))
		}
		return e.StacktraceError(err)
	}
	return nil
}
