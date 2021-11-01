package route

import (
	"net/http"
	"nuxt-echo-chat/backend/internal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	route := echo.New()
	route.Use(middleware.Logger())
	route.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!!")
	})

	g := route.Group("/api")

	internal.NewMelody()
	g.GET("/chat/:roomid", internal.ChatWebsocket)
	g.POST("/fetchChat/:roomid", internal.FetchChat)
	internal.DefineMelodyBehavior()
	return route
}