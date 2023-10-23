package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	chat "main/microservices/chatServer/gen_files"

	conf "main/config"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	myRouter := mux.NewRouter()

	urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	//urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
	config, _ := pgxpool.ParseConfig(urlDB)
	config.MaxConns = 70
	db, err := pgxpool.New(context.Background(), config.ConnString())

	if err != nil {
		log.Println("could not connect to database")
	} else {
		log.Println("database is reachable")
	}
	defer db.Close()

	hub := chat.NewHub()
	go hub.Run()
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { log.Println("main page") })
	myRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { chat.ServeWs(hub, w, r) })
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	chat.ServeWs(hub, w, r)
	// })

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Println("cant listen grpc port", err)
	}
	server := grpc.NewServer()
	chat.RegisterBotChatServer(server, NewChatManager(db, hub))
	log.Println("starting grpc server at :8082")
	go server.Serve(lis)

	log.Println("starting web server at :8081")
	err = http.ListenAndServe(":8081", myRouter)

	if err != nil {
		log.Println("cant serve", err)
	}

}

const sessKeyLen = 10

type ChatManager struct {
	chat.UnimplementedBotChatServer
	db  *pgxpool.Pool
	mu  sync.RWMutex
	hub *chat.Hub
}

func NewChatManager(db *pgxpool.Pool, hub *chat.Hub) *ChatManager {
	return &ChatManager{
		mu:  sync.RWMutex{},
		db:  db,
		hub: hub,
	}
}

func (sm *ChatManager) Recieve(ctx context.Context, in *chat.Message) (*chat.Status, error) {
	log.Println("call Receive ", in)
	// req, err := json.Marshal(in)
	// if err != nil {
	// 	log.Println(err)
	// 	return &chat.Status{IsSuccessful: false}, nil
	// }
	//sm.hub.Broadcast <- []byte(req)
	mes := chat.MessageWebsocket{Text: in.Text, ChatID: in.ChatID}
	log.Println(mes)
	sm.hub.Broadcast <- &mes
	// _, err := sm.db.Query(context.Background(), `INSERT INTO messages (chatID, text, isAuthorTeacher, time) VALUES ($1, $2, $3, $4);`, in.ChatID, in.Text, false, time.Now().Format("2006.01.02 15:04:05"))
	// if err != nil {
	// 	log.Println(err)
	// 	return &chat.Status{IsSuccessful: false}, nil
	// }

	return &chat.Status{IsSuccessful: true}, nil
}
