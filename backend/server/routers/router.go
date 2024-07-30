package routers

import (
	"net/http"
	"time"

	"github.com/TimurZheksimbaev/Golang-webchat/config"
	"github.com/TimurZheksimbaev/Golang-webchat/server/user"
	"github.com/TimurZheksimbaev/Golang-webchat/server/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func InitRouter(appConfig *config.AppConfig, userHandler *user.Handler, wsHandler *ws.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{appConfig.FrontendURL},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func (origin string) bool  {
			return origin == appConfig.FrontendURL
		},
		MaxAge: 12 * time.Hour,
	}))

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