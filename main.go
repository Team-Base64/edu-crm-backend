package main

import (
	"context"
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
	"github.com/jackc/pgx/v5/pgxpool"

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
	config, _ := pgxpool.ParseConfig(os.Getenv(conf.UrlDB))
	// urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	// config, _ := pgxpool.ParseConfig(urlDB)

	config.MaxConns = 70
	db, err := pgxpool.New(context.Background(), config.ConnString())

	if err != nil {
		log.Println("could not connect to database")
	} else {
		log.Println("database is reachable")
	}
	defer db.Close()

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
