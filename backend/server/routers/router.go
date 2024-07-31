package routers

import (
	"net/http"
	"github.com/TimurZheksimbaev/Golang-webchat/server/user"
	"github.com/TimurZheksimbaev/Golang-webchat/server/ws"
	"github.com/gin-gonic/gin"
)



func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("", func (ctx *gin.Context)  {
		ctx.String(http.StatusOK, "Welcome home")
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Page not found")
	})

	auth := r.Group("/auth")
	websocket := r.Group("/ws")

	auth.POST("/register", userHandler.CreateUser)
	auth.POST("/login", userHandler.Login)
	auth.GET("/logout", userHandler.Logout)

	websocket.POST("/createRoom", wsHandler.CreateRoom)
	websocket.GET("/joinRoom/:roomId", wsHandler.JoinRoom)
	websocket.GET("/rooms", wsHandler.GetRooms)
	websocket.GET("/clients/:roomId", wsHandler.GetClients)

	return r
}