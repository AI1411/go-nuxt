package routes

import (
	"github.com/labstack/echo"
	"go-nuxt-youtube/web/api"
)

func Init(e *echo.Echo)  {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
		g.GET("/video/:id", api.GetVideo())
	}
}