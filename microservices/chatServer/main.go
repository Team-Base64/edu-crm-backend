package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"main/domain/model"
	chat "main/microservices/chatServer/gen_files"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	myRouter := mux.NewRouter()

	// urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	urlDB := os.Getenv("URL_DB")
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

func (sm *ChatManager) AddMessage(in *model.CreateMessage) error {
	_, err := sm.db.Query(context.Background(), `INSERT INTO messages (chatID, text, isAuthorTeacher, time) VALUES ($1, $2, $3, $5);`, in.ChatID, in.Text, in.IsAuthorTeacher, time.Now().Format("2006.01.02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}

func (sm *ChatManager) StartChatTG(ch chat.BotChat_StartChatTGServer) error {
	log.Println("start chat tg")
	go func() {
		for {
			// отправка из вебсокета в бота
			mes2 := <-sm.hub.MessagesToTGBot
			mockChatID := 1
			resp := chat.Message{Text: mes2.Text, ChatID: int32(mockChatID)}
			if err := ch.Send(&resp); err != nil {
				log.Println(err)
				if err.Error() == "rpc error: code = Canceled desc = context canceled" {
					log.Println("breaking grpc stream")
					break
					//return nil
				}
				continue
			}
			sm.AddMessage(&model.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true})
		}
	}()
	for {
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			if err.Error() == "rpc error: code = Canceled desc = context canceled" {
				log.Println("breaking grpc stream")
				return nil
			}
			continue
		}
		log.Println(req)
		mes := chat.MessageWebsocket{Text: req.Text, ChatID: req.ChatID, Channel: "chat"}
		sm.hub.Broadcast <- &mes
		sm.AddMessage(&model.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false})
	}
}

func (sm *ChatManager) StartChatVK(ch chat.BotChat_StartChatVKServer) error {
	log.Println("start chat vk")
	go func() {
		for {
			// отправка из вебсокета в бота
			mes2 := <-sm.hub.MessagesToVKBot
			resp := chat.Message{Text: mes2.Text, ChatID: 1}
			if err := ch.Send(&resp); err != nil {
				log.Println(err)
				if err.Error() == "rpc error: code = Canceled desc = context canceled" {
					log.Println("breaking grpc stream")
					break
					//return nil
				}
				continue
			}
			sm.AddMessage(&model.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true})
		}
	}()
	for {
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			if err.Error() == "rpc error: code = Canceled desc = context canceled" {
				log.Println("breaking grpc stream")
				return nil
			}
			continue
		}
		log.Println(req)
		mes := chat.MessageWebsocket{Text: req.Text, ChatID: req.ChatID, Channel: "chat"}
		sm.hub.Broadcast <- &mes
		sm.AddMessage(&model.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false})
	}
}
