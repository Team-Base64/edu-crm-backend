package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	grpcCalendar "main/delivery/grpc/calendar"
	protoCalendar "main/delivery/grpc/calendar/proto"
	grpcChat "main/delivery/grpc/chat"
	protoChat "main/delivery/grpc/chat/proto"
	handler "main/delivery/http"
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
var calendarGrpcUrl string
var tokenLen int
var tokenLetters string
var tokenFile string
var credentialsFile string
var filestoragePath string
var urlDomain string

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

	calendarGrpcUrl, exist = os.LookupEnv(conf.CalendarGrpcUrl)
	if !exist || len(calendarGrpcUrl) == 0 {
		log.Fatalln("could not get calendar grpc url from env")
	}

	// pgUser, exist := os.LookupEnv(conf.PG_USER)
	// if !exist || len(pgUser) == 0 {
	// 	log.Fatalln("could not get database host from env")
	// }
	// pgPwd, exist := os.LookupEnv(conf.PG_PWD)
	// if !exist || len(pgPwd) == 0 {
	// 	log.Fatalln("could not get database password from env")
	// }
	// pgHost, exist := os.LookupEnv(conf.PG_HOST)
	// if !exist || len(pgHost) == 0 {
	// 	log.Fatalln("could not get database host from env")
	// }
	// pgPort, exist := os.LookupEnv(conf.PG_PORT)
	// if !exist || len(pgPort) == 0 {
	// 	log.Fatalln("could not get database port from env")
	// }
	// pgDB, exist := os.LookupEnv(conf.PG_DB)
	// if !exist || len(pgDB) == 0 {
	// 	log.Fatalln("could not get database name from env")
	// }
	// urlDB = "postgres://" + pgUser + ":" + pgPwd + "@" + pgHost + ":" + pgPort + "/" + pgDB

	urlDBs, exist := os.LookupEnv(conf.URL_DB)
	if !exist || len(urlDBs) == 0 {
		log.Fatalln("could not get database name from env")
	}
	urlDB = urlDBs

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

	filestoragePath, exist = os.LookupEnv(conf.FilestoragePath)
	if !exist || len(filestoragePath) == 0 {
		log.Fatalln("could not get filestorage path from env")
	}

	urlDomain, exist = os.LookupEnv(conf.UrlDomain)
	if !exist || len(urlDomain) == 0 {
		log.Fatalln("could not get url domain path from env")
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
	//log.Println(urlDB)
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

	grcpConnCalendar, err := grpc.Dial(
		calendarGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("cant create connecter to grpc calendar")
	}
	log.Println("connecter to grpc calendar service is created")
	defer grcpConnCalendar.Close()

	chatService := grpcChat.NewChatService(protoChat.NewChatControllerClient(grcpConnChat))
	calendarService := grpcCalendar.NewCalendarService(protoCalendar.NewCalendarControllerClient(grcpConnCalendar))

	Usecase := usecase.NewUsecase(
		Store,
		tokenLetters,
		tokenLen,
		chatService,
		calendarService,
		tokenFile,
		credentialsFile,
	)

	Handler := handler.NewHandler(Usecase, filestoragePath, urlDomain)

	myRouter.HandleFunc(conf.PathAttach, Handler.UploadFile).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathLogin, Handler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogout, Handler.Logout).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathAuth, Handler.CheckAuth).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathSignUp, Handler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, Handler.GetTeacherProfile).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathChats, Handler.GetTeacherChats).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathChatByID, Handler.GetChat).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathChatByID, Handler.ReadChat).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathCalendar, Handler.GetCalendar).Methods(http.MethodGet, http.MethodOptions)
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

	myRouter.HandleFunc(conf.PathTasks, Handler.GetTeacherTasks).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathTasks, Handler.CreateTask).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathTaskByID, Handler.GetTaskByID).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathSolution, Handler.GetSolution).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSolution, Handler.AddEvaluationForSolution).Methods(http.MethodPut, http.MethodOptions)

	myRouter.HandleFunc(conf.PathStudent, Handler.GetStudent).Methods(http.MethodGet, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	amw := handler.NewAuthMiddleware(Usecase)
	myRouter.Use(amw.CheckAuthMiddleware)

	err = http.ListenAndServe(conf.Port, myRouter)
	if err != nil {
		log.Println(e.StacktraceError(err))
	}
}
