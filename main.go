package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	ctrl "main/controller"
	"main/delivery"
	"main/repository"
	"main/usecase"

	conf "main/config"
	_ "main/docs"
	e "main/domain/errors"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	httpSwagger "github.com/swaggo/http-swagger"
)

var urlDB string
var chatGrpcUrl string
var tokenLen int
var tokenLetters string
var tokenFile string
var credentialsFile string
var baseFilestorage string

func loggingAndCORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}

		next.ServeHTTP(w, r)
	})
}

func init() {
	var err error
	var exist bool

	chatGrpcUrl, exist = os.LookupEnv(conf.ChatGrpcUrl)
	if !exist || len(chatGrpcUrl) == 0 {
		log.Fatalln("could not get chat grpc url from env")
	}

	urlDB, exist = os.LookupEnv(conf.UrlDB)
	if !exist || len(urlDB) == 0 {
		log.Fatalln("could not get database url from env")
	}

	tokenLen, err = strconv.Atoi(os.Getenv(conf.TokenLenght))
	if err != nil {
		log.Fatalln("could not get token length from env")
	}

	tokenLetters, exist = os.LookupEnv(conf.TokenLetters)
	if !exist || len(tokenLetters) == 0 {
		log.Fatalln("could not get token letters from env")
	}

	tokenFile, exist = os.LookupEnv(conf.TokenFile)
	if !exist || len(tokenFile) == 0 {
		log.Fatalln("could not get token file path from env")
	}

	credentialsFile, exist = os.LookupEnv(conf.CredentialsFile)
	if !exist || len(credentialsFile) == 0 {
		log.Fatalln("could not get credentials file path from env")
	}

	baseFilestorage, exist = os.LookupEnv(conf.BaseFilestorage)
	if !exist || len(baseFilestorage) == 0 {
		log.Fatalln("could not get base filestorage path from env")
	}
}

func main() {
	log.SetFlags(log.LstdFlags)

	myRouter := mux.NewRouter()

	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Fatalln("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln("unable to reach database ", err)
	}
	log.Println("database is reachable")

	Store := repository.NewStore(db)

	grcpConnChat, err := grpc.Dial(
		chatGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("cant create connecter to grpc chat")
	}
	log.Println("connecter to grpc chat service is created")
	defer grcpConnChat.Close()

	ChatService := ctrl.NewChatService(ctrl.NewBotChatClient(grcpConnChat))

	Usecase := usecase.NewUsecase(
		Store,
		tokenLetters,
		tokenLen,
		ChatService,
		tokenFile,
		credentialsFile,
	)

	Handler := delivery.NewHandler(Usecase, baseFilestorage)

	myRouter.HandleFunc(conf.PathAttach, Handler.UploadFile).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathOAuthSetToken, Handler.SetOAUTH2Token).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathOAuthSaveToken, Handler.SaveOAUTH2TokenToFile).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathSignUp, Handler.CreateTeacher).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, Handler.GetTeacherProfile).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathChats, Handler.GetTeacherChats).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathChatByID, Handler.GetChat).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathCalendar, Handler.CreateCalendar).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathAddEvent, Handler.CreateCalendarEvent).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathGetEvents, Handler.GetCalendarEvents).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathEvent, Handler.DeleteCalendarEvent).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathEvent, Handler.UpdateCalendarEvent).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathClasses, Handler.GetTeacherClasses).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassByID, Handler.GetClass).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClasses, Handler.CreateClass).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassStudents, Handler.GetStudentsFromClass).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassFeed, Handler.GetClassFeed).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassFeed, Handler.CreatePost).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassHomeworks, Handler.GetHomeworksFromClass).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassSolutions, Handler.GetSolutionsFromClass).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathHomeworkByID, Handler.GetHomework).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathHomework, Handler.CreateHomework).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathHomeworkSolutions, Handler.GetSolutionsForHomework).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathSolution, Handler.GetSolution).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathStudent, Handler.GetStudent).Methods(http.MethodGet, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	err = http.ListenAndServe(conf.Port, myRouter)
	if err != nil {
		log.Println(e.StacktraceError(err))
	}
}
