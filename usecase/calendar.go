package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(tokFile string, config *oauth2.Config) (*http.Client, error) {
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	tokenSource := config.TokenSource(context.Background(), tok)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	if newToken.AccessToken != tok.AccessToken {
		saveToken(tokFile, newToken)
		log.Println("Saved new token:", newToken.AccessToken)
	}
	return config.Client(context.Background(), tok), nil
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	log.Println("Saving credential file to: ", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Println("Unable to cache oauth token: ", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func (uc *Usecase) getCalendarServicelient() (*calendar.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile(uc.credentialsFile)
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return nil, e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return nil, e.StacktraceError(err)
	}
	client, err := getClient(uc.tokenFile, config)
	if err != nil {
		log.Println("Unable to get client from token: ", err)
		return nil, e.StacktraceError(err)
	}

	return calendar.NewService(ctx, option.WithHTTPClient(client))

}

func (uc *Usecase) SetOAUTH2Token() error {
	//ctx := context.Background()
	b, err := os.ReadFile(uc.credentialsFile)
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return e.StacktraceError(err)
	}
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	//browser.OpenURL(authURL)
	exec.Command("xdg-open", authURL).Start()
	return nil
	// client, err := getClient(config)
	// if err != nil {
	// 	log.Println("Unable to get client from token: ", err)
	// 	return  e.StacktraceError(err)
	// }
}

func (uc *Usecase) SaveOAUTH2Token(authCode string) error {
	//ctx := context.Background()
	b, err := os.ReadFile(uc.credentialsFile)
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return e.StacktraceError(err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Println("Unable to retrieve token from web: ", err)
		return e.StacktraceError(err)
	}

	saveToken(uc.tokenFile, token)
	return nil
}

func (uc *Usecase) CreateCalendar(teacherID int) (*model.CreateCalendarResponse, error) {

	srv, err := uc.getCalendarServicelient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return nil, e.StacktraceError(err)
	}

	newCal := &calendar.Calendar{TimeZone: "Europe/Moscow", Summary: "EDUCRM Calendar"}
	cal, err := srv.Calendars.Insert(newCal).Do()
	if err != nil {
		log.Println("Unable to create calendar: ", err)
		return nil, e.StacktraceError(err)
	}

	Acl := &calendar.AclRule{Scope: &calendar.AclRuleScope{Type: "default"}, Role: "reader"}
	_, err = srv.Acl.Insert(cal.Id, Acl).Do()
	if err != nil {
		log.Println("Unable to create ACL: ", err)
		return nil, e.StacktraceError(err)
	}

	innerID, err := uc.store.CreateCalendarDB(teacherID, cal.Id)
	if err != nil {
		log.Println("DB err: ", err)
		return nil, e.StacktraceError(err)
	}

	return &model.CreateCalendarResponse{ID: innerID, IDInGoogle: cal.Id}, nil
}

func (uc *Usecase) CreateCalendarEvent(req *model.CalendarEvent, teacherID int) error {
	srv, err := uc.getCalendarServicelient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return e.StacktraceError(err)
	}

	event := &calendar.Event{
		Summary:     req.Title + " Class " + fmt.Sprintf("%d", req.ClassID),
		Description: req.Description,
		Start: &calendar.EventDateTime{
			DateTime: req.StartDate,
			//TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: req.EndDate,
			//TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}
	calendarID, err := uc.store.GetCalendarGoogleID(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}

	_, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Println("Unable to create event: ", err)
		return e.StacktraceError(err)
	}

	bcMsg := model.ClassBroadcastMessage{
		ClassID:     req.ClassID,
		Title:       "Новое событие!" + "\n" + req.Title,
		Description: req.Description + "\n" + "Начало: " + req.StartDate + "\n" + "Окончание: " + req.EndDate,
		Attaches:    []string{},
	}
	if err := uc.chatService.BroadcastMsg(&bcMsg); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (uc *Usecase) GetCalendarEvents(teacherID int) ([]model.CalendarEvent, error) {
	srv, err := uc.getCalendarServicelient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return nil, e.StacktraceError(err)
	}
	calendarID, err := uc.store.GetCalendarGoogleID(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return nil, e.StacktraceError(err)
	}
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(calendarID).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(100).OrderBy("startTime").Do()
	if err != nil {
		log.Println("Unable to retrieve next ten of the user's events: ", err)
		return nil, e.StacktraceError(err)
	}

	ans := []model.CalendarEvent{}
	for _, item := range events.Items {
		s := strings.Split(item.Summary, " ")
		classID := 0
		if len(s) > 2 && s[len(s)-2] == "Class" {
			classIDs := s[len(s)-1]
			classID, err = strconv.Atoi(classIDs)
			if err != nil {
				log.Println("error: ", err)
			}
		}

		tmp := model.CalendarEvent{Title: item.Summary, Description: item.Description,
			StartDate: item.Start.DateTime, EndDate: item.End.DateTime, ClassID: classID, ID: item.Id}
		ans = append(ans, tmp)
	}

	return ans, nil
}

func (uc *Usecase) DeleteCalendarEvent(teacherID int, eventID string) error {
	srv, err := uc.getCalendarServicelient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return e.StacktraceError(err)
	}
	calendarID, err := uc.store.GetCalendarGoogleID(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}
	err = srv.Events.Delete(calendarID, eventID).Do()
	if err != nil {
		log.Println("Unable to delete event: ", err)
		return e.StacktraceError(err)
	}

	return nil
}

func (uc *Usecase) UpdateCalendarEvent(req *model.CalendarEvent, teacherID int) error {
	srv, err := uc.getCalendarServicelient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return e.StacktraceError(err)
	}
	s := strings.Split(req.Title, " ")
	//log.Println(s)
	newTitle := ""
	if len(s) > 2 && s[len(s)-2] == "Class" {
		newTitle = strings.Join(s[:len(s)-2], " ")
	} else {
		newTitle = req.Title
	}

	event := &calendar.Event{
		Summary:     newTitle + " Class " + fmt.Sprintf("%d", req.ClassID),
		Description: req.Description,
		Start: &calendar.EventDateTime{
			DateTime: req.StartDate,
			//TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: req.EndDate,
			//TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}
	calendarID, err := uc.store.GetCalendarGoogleID(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}

	_, err = srv.Events.Update(calendarID, req.ID, event).Do()
	if err != nil {
		log.Println("Unable to update event: ", err)
		return e.StacktraceError(err)
	}

	bcMsg := model.ClassBroadcastMessage{
		ClassID:     req.ClassID,
		Title:       "Событие обновлено!" + "\n" + req.Title,
		Description: req.Description + "\n" + "Начало: " + req.StartDate + "\n" + "Окончание: " + req.EndDate,
		Attaches:    []string{},
	}
	if err := uc.chatService.BroadcastMsg(&bcMsg); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}
