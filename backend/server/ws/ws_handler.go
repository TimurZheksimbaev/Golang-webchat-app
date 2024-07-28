package ws

import (
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	WRITE_BUFFER_SIZE = 1024
	READ_BUFFER_SIZE = 1024
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	h.hub.Rooms[req.ID] = &Room{
		ID: req.ID,
		Name: req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: READ_BUFFER_SIZE,
	WriteBufferSize: WRITE_BUFFER_SIZE, 
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn: conn,
		Message: make(chan *Message, 10),
		ID: clientID,
		RoomID: roomID,
		Username: username,
	}

	m := &Message{
		Content: "A new user has joined the room",
		RoomID: roomID,
		Username: username,
	}

	// Register a new client
	h.hub.Register <- cl

	// Broadcast client's message
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}


func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomResponse, 0)
	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID: r.ID,
			Name: r.Name,
		})
	}
	c.JSON(http.StatusOK, rooms)
}


func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientResponse
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientResponse, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientResponse{
			ID: c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}