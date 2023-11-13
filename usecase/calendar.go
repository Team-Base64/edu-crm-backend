package usecase

import (
	"context"
	"encoding/json"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) (*http.Client, error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	//now := time.Now()
	//redTime := now.Add(1 * time.Minute)
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	tokenSource := config.TokenSource(context.Background(), tok)
	//tokenSource := conf.TokenSource(oauth2.NoContext, token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	if newToken.AccessToken != tok.AccessToken {
		saveToken(tokFile, newToken)
		log.Println("Saved new token:", newToken.AccessToken)
	}

	//if err != nil {
	//log.Println(err)
	// if tok.Expiry.Before(redTime) {
	// 	exec.Command("xdg-open", "http://127.0.0.1:8080/api/oauth").Start()
	// }
	if err != nil {
		//tok = getTokenFromWeb(config)
		//saveToken(tokFile, tok)
		return nil, e.StacktraceError(err)
		//return config.Client(context.Background(), tok), nil
	}
	return config.Client(context.Background(), tok), nil
}

// // Request a token from the web, then returns the retrieved token.
// func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	// fmt.Printf("Go to the following link in your browser then type the "+
// 	// 	"authorization code: \n%v\n", authURL)
// 	//открыть новую вкладку
// 	//exec.Command("xdg-open", authURL).Start()

// 	log.Println("authorization code: ")
// 	browser.OpenURL(authURL)
// 	var authCode string
// 	if _, err := fmt.Scan(&authCode); err != nil {
// 		log.Fatalf("Unable to read authorization code: %v", err)
// 	}

// 	tok, err := config.Exchange(context.TODO(), authCode)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve token from web: %v", err)
// 	}
// 	//tok := "0AfJohXk2EWvIKvCXEK_vUSJRn7W4eoy2RqQX-wWwiYavt7Q0KOi6nnAlOb_tKn_7IM0tlg"
// 	return tok
// }

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

func (uc *Usecase) SetOAUTH2Token() error {
	//ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
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
	b, err := os.ReadFile("credentials.json")
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

	path := "token.json"
	saveToken(path, token)
	// log.Println("Saving credential file to: ", path)
	// f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	// if err != nil {
	// 	log.Println("Unable to cache oauth token: ", err)
	// 	return e.StacktraceError(err)
	// }
	// defer f.Close()
	// json.NewEncoder(f).Encode(token)
	return nil
	//authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	//browser.OpenURL(authURL)
	//return nil
	// client, err := getClient(config)
	// if err != nil {
	// 	log.Println("Unable to get client from token: ", err)
	// 	return  e.StacktraceError(err)
	// }
}

func (uc *Usecase) CreateCalendar(teacherID int) (*model.CreateCalendarResponse, error) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return nil, e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return nil, e.StacktraceError(err)
	}

	client, err := getClient(config)
	if err != nil {
		log.Println("Unable to get client from token: ", err)
		return nil, e.StacktraceError(err)
	}

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
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

func (uc *Usecase) CreateCalendarEvent(req *model.CreateCalendarEvent, teacherID int, classID int) error {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return e.StacktraceError(err)
	}

	client, err := getClient(config)
	if err != nil {
		log.Println("Unable to get client from token: ", err)
		return e.StacktraceError(err)
	}

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return e.StacktraceError(err)
	}

	event := &calendar.Event{
		Summary:     req.Title,
		Description: req.Description,
		Start: &calendar.EventDateTime{
			DateTime: req.StartDate,
			TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: req.EndDate,
			TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}
	calendarID, err := uc.store.GetCalendarGoogleID(teacherID)
	if err != nil {
		log.Println("DB err: ", err)
		return e.StacktraceError(err)
	}
	//log.Println(calendarID, req.StartDate, req.EndDate)
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Println("Unable to create event: ", err)
		return e.StacktraceError(err)
	}

	//MockClassID := 1
	bcMsg := model.ClassBroadcastMessage{
		ClassID:     classID,
		Title:       "Новое событие!" + "\n" + req.Title,
		Description: req.Description + "\n" + "Начало: " + req.StartDate + "\n" + "Окончание: " + req.EndDate,
		Attaches:    []string{},
	}
	if err := uc.chatService.BroadcastMsg(&bcMsg); err != nil {
		return err
	}
	return nil
}
