package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	conf "main/config"
	grpcCalendar "main/delivery/grpc/calendar"
	protoCalendar "main/delivery/grpc/calendar/proto"
	grpcChat "main/delivery/grpc/chat"
	protoChat "main/delivery/grpc/chat/proto"
	httpHandler "main/delivery/http"
	e "main/domain/errors"
	localStore "main/repository/local-storage"
	pgStore "main/repository/pg"
	backUsecase "main/usecase/backend"

	_ "main/docs"

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
var filestoragePath string
var chatFilesPath string
var homeworkFilesPath string
var solutionFilesPath string
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

	chatGrpcUrl, exist = os.LookupEnv(conf.CHAT_GRPC_URL)
	if !exist || len(chatGrpcUrl) == 0 {
		log.Fatalln("could not get chat grpc url from env")
	}

	calendarGrpcUrl, exist = os.LookupEnv(conf.CALENDAR_GRPC_URL)
	if !exist || len(calendarGrpcUrl) == 0 {
		log.Fatalln("could not get calendar grpc url from env")
	}

	pgUser, exist := os.LookupEnv(conf.PG_USER)
	if !exist || len(pgUser) == 0 {
		log.Fatalln("could not get database host from env")
	}

	pgPwd, exist := os.LookupEnv(conf.PG_PWD)
	if !exist || len(pgPwd) == 0 {
		log.Fatalln("could not get database password from env")
	}

	pgHost, exist := os.LookupEnv(conf.PG_HOST)
	if !exist || len(pgHost) == 0 {
		log.Fatalln("could not get database host from env")
	}

	pgPort, exist := os.LookupEnv(conf.PG_PORT)
	if !exist || len(pgPort) == 0 {
		log.Fatalln("could not get database port from env")
	}

	pgDB, exist := os.LookupEnv(conf.PG_DB)
	if !exist || len(pgDB) == 0 {
		log.Fatalln("could not get database name from env")
	}

	urlDB = "postgres://" + pgUser + ":" + pgPwd + "@" + pgHost + ":" + pgPort + "/" + pgDB

	tokenLen, err = strconv.Atoi(os.Getenv(conf.TOKEN_LENGTH))
	if err != nil {
		log.Fatalln("could not get token length from env")
	}

	tokenLetters, exist = os.LookupEnv(conf.TOKEN_LETTERS)
	if !exist || len(tokenLetters) == 0 {
		log.Fatalln("could not get token letters from env")
	}

	filestoragePath, exist = os.LookupEnv(conf.FILESTORAGE_PATH)
	if !exist || len(filestoragePath) == 0 {
		log.Fatalln("could not get filestorage path from env")
	}

	chatFilesPath, exist = os.LookupEnv(conf.CHAT_FILES_PATH)
	if !exist || len(chatFilesPath) == 0 {
		log.Fatalln("could not get chat files path from env")
	}

	homeworkFilesPath, exist = os.LookupEnv(conf.HOMEWORK_FILES_PATH)
	if !exist || len(homeworkFilesPath) == 0 {
		log.Fatalln("could not get homework files path from env")
	}

	solutionFilesPath, exist = os.LookupEnv(conf.SOLUTION_FILES_PATH)
	if !exist || len(solutionFilesPath) == 0 {
		log.Fatalln("could not get solution files path from env")
	}

	urlDomain, exist = os.LookupEnv(conf.URL_DOMAIN)
	if !exist || len(urlDomain) == 0 {
		log.Fatalln("could not get url domain path from env")
	}
}

func main() {
	log.SetFlags(log.LstdFlags)

	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Fatalln("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln("unable to reach database ", err)
	}
	log.Println("database is reachable")

	dataStore := pgStore.NewPostgreSqlStore(db)

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

	chatService := grpcChat.NewChatService(protoChat.NewChatClient(grcpConnChat))
	calendarService := grpcCalendar.NewCalendarService(protoCalendar.NewCalendarClient(grcpConnCalendar))

	fileStore := localStore.NewLocalStore(
		chatFilesPath,
		homeworkFilesPath,
		solutionFilesPath,
		filestoragePath,
	)

	usecase := backUsecase.NewBackendUsecase(
		dataStore,
		tokenLetters,
		tokenLen,
		chatService,
		calendarService,
		fileStore,
		urlDomain,
	)

	myRouter := mux.NewRouter()
	handler := httpHandler.NewHandler(usecase)

	myRouter.HandleFunc(conf.PathAttach, handler.UploadFile).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathLogin, handler.Login).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathLogout, handler.Logout).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathAuth, handler.CheckAuth).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathSignUp, handler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, handler.GetTeacherProfile).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathChats, handler.GetTeacherChats).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathChatByID, handler.GetChat).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathChatByID, handler.ReadChat).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathCalendar, handler.GetCalendar).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathCalendar, handler.CreateCalendar).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathAddEvent, handler.CreateCalendarEvent).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathGetEvents, handler.GetCalendarEvents).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathEvent, handler.DeleteCalendarEvent).Methods(http.MethodDelete, http.MethodOptions)
	myRouter.HandleFunc(conf.PathEvent, handler.UpdateCalendarEvent).Methods(http.MethodPost, http.MethodOptions)

	myRouter.HandleFunc(conf.PathClasses, handler.GetTeacherClasses).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassByID, handler.GetClass).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClasses, handler.CreateClass).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassStudents, handler.GetStudentsFromClass).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassFeed, handler.GetClassFeed).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassFeed, handler.CreatePost).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassHomeworks, handler.GetHomeworksFromClass).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathClassSolutions, handler.GetSolutionsFromClass).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathHomeworkByID, handler.GetHomework).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathHomework, handler.CreateHomework).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathHomeworkSolutions, handler.GetSolutionsForHomework).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathTasks, handler.GetTeacherTasks).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathTasks, handler.CreateTask).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathTaskByID, handler.GetTaskByID).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathSolution, handler.GetSolution).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathSolution, handler.AddEvaluationForSolution).Methods(http.MethodPut, http.MethodOptions)

	myRouter.HandleFunc(conf.PathStudent, handler.GetStudent).Methods(http.MethodGet, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	amw := httpHandler.NewAuthMiddleware(usecase)
	myRouter.Use(amw.CheckAuthMiddleware)

	err = http.ListenAndServe(conf.Port, myRouter)
	if err != nil {
		log.Println(e.StacktraceError(err))
	}
}
