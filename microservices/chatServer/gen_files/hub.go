package chat

type MessageWebsocket struct {
	Text   string `json:"text"`
	ChatID int32  `json:"chatid"`
}

// type Connection struct {
// 	ws   *websocket.Conn
// 	send chan MessageWebsocket
// }

type Hub struct {
	clients         map[*Client]bool
	Broadcast       chan *MessageWebsocket
	MessagesToTGBot chan *MessageWebsocket
	register        chan *Client
	unregister      chan *Client
	chats           map[int32]*Client // соединение по id чата
	clientChats     map[*Client][]int32
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:       make(chan *MessageWebsocket),
		MessagesToTGBot: make(chan *MessageWebsocket),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		clientChats:     make(map[*Client][]int32),
		clients:         make(map[*Client]bool),
		chats:           make(map[int32]*Client),
		//Messages:   make(chan *MessageWebsocket),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				for _, cl := range h.clientChats[client] {
					delete(h.chats, cl)
				}
				delete(h.clientChats, client)
				delete(h.clients, client)
				close(client.send)
			}
			// case message := <-h.Broadcast:
			// 	for client := range h.clients {
			// 		select {
			// 		case client.send <- message:
			// 		default:
			// 			close(client.send)
			// 			delete(h.clients, client)
			// 		}
			// 	}
		case mes := <-h.Broadcast:
			// for _, client := range h.clients {
			// 	log.Println(client, "  ", h.chats[2], mes.ChatID, h.chats[mes.ChatID])
			// }
			conn := h.chats[mes.ChatID]
			if conn == nil {
				break
			}
			//log.Println(conn)
			select {
			case conn.send <- mes:
			default:
				close(conn.send)
				delete(h.chats, mes.ChatID)
				// delete(chats, c)
				// if len(chats) == 0 {
				// 	delete(h.chats, mes.ChatID)
				// }
			}

		}
	}
}
