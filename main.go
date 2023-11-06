package main

import (
	"context"
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

	"github.com/jackc/pgx/v5"

	httpSwagger "github.com/swaggo/http-swagger"
)

func loggingAndCORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.SetFlags(log.LstdFlags)

	myRouter := mux.NewRouter()

	db, err := pgx.Connect(context.Background(), os.Getenv(conf.UrlDB))
	if err != nil {
		log.Fatalln("could not connect to database")
	}
	defer db.Close(context.Background())

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalln("unable to reach database ", err)
	}
	log.Println("database is reachable")

	Store := repository.NewStore(db)

	tokenLen, err := strconv.Atoi(os.Getenv(conf.TokenLenght))
	if err != nil {
		log.Fatalln("could not get token length from env")
	}
	tokenLetters, exist := os.LookupEnv(conf.TokenLetters)
	if !exist || len(tokenLetters) == 0 {
		log.Fatalln("could not get token letters from env")
	}

	grcpConnChat, err := grpc.Dial(
		os.Getenv(conf.ChatGrpcUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("cant connect to grpc chat")
	} else {
		log.Println("connected to grpc chat service")
	}
	defer grcpConnChat.Close()

	ChatService := ctrl.NewChatService(ctrl.NewBotChatClient(grcpConnChat))

	Usecase := usecase.NewUsecase(
		Store,
		tokenLetters,
		tokenLen,
		ChatService,
	)

	Handler := delivery.NewHandler(Usecase)

	myRouter.HandleFunc(conf.PathSignUp, Handler.CreateTeacher).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, Handler.GetTeacherProfile).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathChats, Handler.GetTeacherChats).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathChatByID, Handler.GetChat).Methods(http.MethodGet, http.MethodOptions)

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

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	err = http.ListenAndServe(conf.Port, myRouter)
	if err != nil {
		log.Println(e.StacktraceError(err))
	}
}
