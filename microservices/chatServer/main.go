package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	chat "main/microservices/chatServer/gen_files"

	conf "main/config"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Println("cant listen port", err)
	}

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

	server := grpc.NewServer()
	chat.RegisterBotChatServer(server, NewChatManager(db))

	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Println("starting server at :8081")
	server.Serve(lis)
}

const sessKeyLen = 10

type ChatManager struct {
	chat.UnimplementedBotChatServer
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewChatManager(db *pgxpool.Pool) *ChatManager {
	return &ChatManager{
		mu: sync.RWMutex{},
		db: db,
	}
}

func (sm *ChatManager) Recieve(ctx context.Context, in *chat.Message) (*chat.Status, error) {
	log.Println("call Receive ", in)
	_, err := sm.db.Query(context.Background(), `INSERT INTO messages (chatID, text, isAuthorTeacher, time) VALUES ($1, $2, $3, $4);`, in.ChatID, in.Text, false, time.Now().Format("2006.01.02 15:04:05"))
	if err != nil {
		log.Println(err)
		return &chat.Status{IsSuccessful: false}, nil
	}

	return &chat.Status{IsSuccessful: true}, nil
}
