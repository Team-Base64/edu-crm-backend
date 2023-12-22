package usecase

import (
	"main/domain/model"
)

func (uc *Usecase) GetCalendar(teacherID int) (*model.CalendarParams, error) {
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

// func (uc *Usecase) CreateCalendarEvent(req *model.CalendarEvent, teacherID int) error {
// 	srv, err := uc.getCalendarServiceClient()
// 	if err != nil {
// 		log.Println("Unable to retrieve calendar Client: ", err)
// 		return e.StacktraceError(err)
// 	}

// 	event := &calendar.Event{
// 		Summary:     req.Title + " Class " + fmt.Sprintf("%d", req.ClassID),
// 		Description: req.Description,
// 		Start: &calendar.EventDateTime{
// 			DateTime: req.StartDate.Format(time.RFC3339Nano),
// 			//TimeZone: "Europe/Moscow",
// 		},
// 		End: &calendar.EventDateTime{
// 			DateTime: req.EndDate.Format(time.RFC3339Nano),
// 			//TimeZone: "Europe/Moscow",
// 		},
// 		Visibility: "public",
// 	}
// 	calendarDB, err := uc.store.GetCalendarDB(teacherID)
// 	if err != nil {
// 		log.Println("DB err: ", err)
// 		return e.StacktraceError(err)
// 	}

// 	event, err = srv.Events.Insert(calendarDB.IDInGoogle, event).Do()

// 	if err != nil {
// 		log.Println("Unable to create event: ", err)
// 		return e.StacktraceError(err)
// 	}

// 	bcMsg := model.ClassBroadcastMessage{
// 		ClassID: req.ClassID,
// 		Title:   "Новое событие!" + "\n" + req.Title,
// 		Description: req.Description + "\n" +
// 			"Начало: " + utils.TimeToString(req.StartDate) + "\n" +
// 			"Окончание: " + utils.TimeToString(req.EndDate) + "\n" +
// 			"Ссылка на календарь: " +
// 			"https://calendar.google.com/calendar/embed?ctz=Europe%2FMoscow&hl=ru&src=" + calendarDB.IDInGoogle,
// 		Attaches: []string{},
// 	}
// 	if err := uc.ctrlService.BroadcastMsg(&bcMsg); err != nil {
// 		err = uc.DeleteCalendarEvent(teacherID, event.Id)
// 		if err != nil {
// 			log.Println("Unable to delete event after bc error: ", err)
// 			//return e.StacktraceError(err)
// 		}
// 		return e.StacktraceError(err)
// 	}
// 	return nil
// }

func (uc *Usecase) GetCalendarEvents(teacherID int) (model.CalendarEvents, error) {

	return uc.ctrlService.GetEvents(teacherID)
}

// func (uc *Usecase) DeleteCalendarEvent(teacherID int, eventID string) error {
// 	srv, err := uc.getCalendarServiceClient()
// 	if err != nil {
// 		log.Println("Unable to retrieve calendar Client: ", err)
// 		return e.StacktraceError(err)
// 	}
// 	calendarDB, err := uc.store.GetCalendarDB(teacherID)
// 	if err != nil {
// 		log.Println("DB err: ", err)
// 		return e.StacktraceError(err)
// 	}
// 	err = srv.Events.Delete(calendarDB.IDInGoogle, eventID).Do()
// 	if err != nil {
// 		log.Println("Unable to delete event: ", err)
// 		return e.StacktraceError(err)
// 	}

// 	return nil
// }

// func (uc *Usecase) UpdateCalendarEvent(req *model.CalendarEvent, teacherID int) error {
// 	srv, err := uc.getCalendarServiceClient()
// 	if err != nil {
// 		log.Println("Unable to retrieve calendar Client: ", err)
// 		return e.StacktraceError(err)
// 	}
// 	s := strings.Split(req.Title, " ")
// 	//log.Println(s)
// 	newTitle := ""
// 	if len(s) > 2 && s[len(s)-2] == "Class" {
// 		newTitle = strings.Join(s[:len(s)-2], " ")
// 	} else {
// 		newTitle = req.Title
// 	}

// 	event := &calendar.Event{
// 		Summary:     newTitle + " Class " + fmt.Sprintf("%d", req.ClassID),
// 		Description: req.Description,
// 		Start: &calendar.EventDateTime{
// 			DateTime: req.StartDate.Format(time.RFC3339Nano),
// 			//TimeZone: "Europe/Moscow",
// 		},
// 		End: &calendar.EventDateTime{
// 			DateTime: req.EndDate.Format(time.RFC3339Nano),
// 			//TimeZone: "Europe/Moscow",
// 		},
// 		Visibility: "public",
// 	}
// 	calendarDB, err := uc.store.GetCalendarDB(teacherID)
// 	if err != nil {
// 		log.Println("DB err: ", err)
// 		return e.StacktraceError(err)
// 	}

// 	event, err = srv.Events.Update(calendarDB.IDInGoogle, req.ID, event).Do()

// 	if err != nil {
// 		log.Println("Unable to update event: ", err)
// 		return e.StacktraceError(err)
// 	}
// 	bcMsg := model.ClassBroadcastMessage{
// 		ClassID: req.ClassID,
// 		Title:   "Событие обновлено!" + "\n" + req.Title,
// 		Description: req.Description + "\n" +
// 			"Начало: " + utils.TimeToString(req.StartDate) + "\n" +
// 			"Окончание: " + utils.TimeToString(req.EndDate) + "\n" +
// 			"Ссылка на календарь: " +
// 			"https://calendar.google.com/calendar/embed?ctz=Europe%2FMoscow&hl=ru&src=" + calendarDB.IDInGoogle,
// 		Attaches: []string{},
// 	}
// 	if err := uc.ctrlService.BroadcastMsg(&bcMsg); err != nil {
// 		err = uc.DeleteCalendarEvent(teacherID, event.Id)
// 		if err != nil {
// 			log.Println("Unable to delete event after bc error: ", err)
// 			//return e.StacktraceError(err)
// 		}
// 		return e.StacktraceError(err)
// 	}
// 	return nil
// }
