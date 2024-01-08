package usecase

import (
	"fmt"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	m "main/domain/model"
	"main/domain/utils"
	"strings"
)

func (uc *Usecase) GetCalendar(teacherID int) (*m.CalendarParams, error) {
	return uc.store.GetCalendarDB(teacherID)
}

// func (uc *Usecase) CreateCalendar(teacherID int) (*model.CalendarParams, error) {

// 	srv, err := uc.getCalendarServiceClient()
// 	if err != nil {
// 		log.Println("Unable to retrieve calendar Client: ", err)
// 		return nil, e.StacktraceError(err)
// 	}

// 	newCal := &calendar.Calendar{TimeZone: "Europe/Moscow", Summary: "EDUCRM Calendar"}
// 	cal, err := srv.Calendars.Insert(newCal).Do()
// 	if err != nil {
// 		log.Println("Unable to create calendar: ", err)
// 		return nil, e.StacktraceError(err)
// 	}

// 	Acl := &calendar.AclRule{Scope: &calendar.AclRuleScope{Type: "default"}, Role: "reader"}
// 	_, err = srv.Acl.Insert(cal.Id, Acl).Do()
// 	if err != nil {
// 		log.Println("Unable to create ACL: ", err)
// 		return nil, e.StacktraceError(err)
// 	}

// 	innerID, err := uc.store.CreateCalendarDB(teacherID, cal.Id)
// 	if err != nil {
// 		log.Println("DB err: ", err)
// 		return nil, e.StacktraceError(err)
// 	}

// 	return &model.CalendarParams{ID: innerID, IDInGoogle: cal.Id}, nil
// }

func (uc *Usecase) CreateCalendarEvent(req *m.CalendarEvent, teacherID int) error {
	calendarDB, err := uc.store.GetCalendarDB(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}
	eventID, err := uc.calendar.CreateEvent(*req, calendarDB.IDInGoogle)
	if err != nil {
		log.Println("Unable to create event: ", err)
		return e.StacktraceError(err)
	}

	bcMsg := model.ClassBroadcastMessage{
		ClassID: req.ClassID,
		Title:   "Новое событие!" + "\n" + req.Title,
		Description: req.Description + "\n" +
			"Начало: " + utils.TimeToString(req.StartDate) + "\n" +
			"Окончание: " + utils.TimeToString(req.EndDate) + "\n" +
			"Ссылка на календарь: " +
			"https://calendar.google.com/calendar/embed?ctz=Europe%2FMoscow&hl=ru&src=" + calendarDB.IDInGoogle,
		Attaches: []string{},
	}
	if err := uc.chat.BroadcastMsg(&bcMsg); err != nil {
		err = uc.calendar.DeleteEvent(calendarDB.IDInGoogle, eventID)
		if err != nil {
			log.Println("Unable to delete event after bc error: ", err)
		}
		log.Println("Unable to broadcast msg: ", err)
		return e.StacktraceError(err)
	}
	return nil
}

func (uc *Usecase) GetCalendarEvents(teacherID int) (model.CalendarEvents, error) {
	return uc.calendar.GetEvents(teacherID)
}

func (uc *Usecase) DeleteCalendarEvent(teacherID int, eventID string) error {
	calendarDB, err := uc.store.GetCalendarDB(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}
	return uc.calendar.DeleteEvent(calendarDB.IDInGoogle, eventID)
}

func (uc *Usecase) UpdateCalendarEvent(req *model.CalendarEvent, teacherID int) error {
	s := strings.Split(req.Title, " ")
	newTitle := ""
	if len(s) > 2 && s[len(s)-2] == "Class" {
		newTitle = strings.Join(s[:len(s)-2], " ")
	} else {
		newTitle = req.Title
	}
	req.Title = newTitle + " Class " + fmt.Sprintf("%d", req.ClassID)

	calendarDB, err := uc.store.GetCalendarDB(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}
	err = uc.calendar.UpdateEvent(*req, calendarDB.IDInGoogle)
	if err != nil {
		log.Println("Unable to create event: ", err)
		return e.StacktraceError(err)
	}

	bcMsg := model.ClassBroadcastMessage{
		ClassID: req.ClassID,
		Title:   "Событие обновлено!" + "\n" + req.Title,
		Description: req.Description + "\n" +
			"Начало: " + utils.TimeToString(req.StartDate) + "\n" +
			"Окончание: " + utils.TimeToString(req.EndDate) + "\n" +
			"Ссылка на календарь: " +
			"https://calendar.google.com/calendar/embed?ctz=Europe%2FMoscow&hl=ru&src=" + calendarDB.IDInGoogle,
		Attaches: []string{},
	}
	if err := uc.chat.BroadcastMsg(&bcMsg); err != nil {
		err = uc.calendar.DeleteEvent(calendarDB.IDInGoogle, req.ID)
		if err != nil {
			log.Println("Unable to delete event after bc error: ", err)
		}
		log.Println("Unable to broadcast msg: ", err)
		return e.StacktraceError(err)
	}
	return nil
}
