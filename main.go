package main

import (
	"log"
	"net/http"
	"os"

	"main/delivery"
	"main/repository"
	"main/usecase"

	conf "main/config"
	_ "main/docs"
	e "main/domain/errors"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"

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

	//urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	//db, err := sql.Open("pgx", urlDB)

	db, err := sql.Open("pgx", os.Getenv(conf.UrlDB))
	if err != nil {
		log.Println("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Println("unable to reach database ", err)
	} else {
		log.Println("database is reachable")
	}

	Store := repository.NewStore(db)

	Usecase := usecase.NewUsecase(Store)

	Handler := delivery.NewHandler(Usecase)

	myRouter.HandleFunc(conf.PathSignUp, Handler.CreateTeacher).Methods(http.MethodPost, http.MethodOptions)
	myRouter.HandleFunc(conf.PathProfile, Handler.GetTeacherProfile).Methods(http.MethodGet, http.MethodOptions)

	myRouter.HandleFunc(conf.PathChats, Handler.GetTeacherChats).Methods(http.MethodGet, http.MethodOptions)
	myRouter.HandleFunc(conf.PathChatByID, Handler.GetChat).Methods(http.MethodGet, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	err = http.ListenAndServe(conf.Port, myRouter)
	if err != nil {
		log.Println(e.StacktraceError(err))
	}
}
