package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	FilmRoutes(e)
	UserRoutes(e)
	AuthRoutes(e)
	CategoryRoutes(e)
	TransactionRoutes(e)
	EpisodeRoutes(e)
}
