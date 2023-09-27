package myserver

import (
	"truck-unii-app/apps/api/handlers"
	"truck-unii-app/apps/api/repositories"
	"truck-unii-app/apps/api/services"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type appRoute struct {
	RoutePath   string
	RouteMethod string
	RouteFunc   func(*fiber.Ctx) error
}

func Route(app *fiber.App, db *gorm.DB) {

	rt := repositories.NewRepoTruck(db)

	st := services.NewServiceTruck(rt)

	h := handlers.NewHandlers()
	_ = handlers.NewHandlerRecycle()
	ht := handlers.NewHandlerTruck(st)

	myroute := []appRoute{}
	myroute = append(myroute, GetRouteTruck(ht)...)
	myroute = append(myroute, GetRouteOther(h)...)
	for _, handle := range myroute {
		log.Info().Msgf("[API][%s] %s", handle.RouteMethod, handle.RoutePath)
		app.Add(handle.RouteMethod, handle.RoutePath, handle.RouteFunc)
	}
}

func GetRouteTruck(h handlers.HandlerTruck) []appRoute {
	return []appRoute{
		{
			RoutePath:   "/truck/profile",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   h.TruckProfileGet,
		},
	}
}

func GetRouteOther(h handlers.Handlers) []appRoute {
	return []appRoute{
		{
			RoutePath:   "/",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   h.Hello,
		},
		{
			RoutePath:   "/swagger/*",
			RouteMethod: fiber.MethodGet,
			RouteFunc:   swagger.HandlerDefault,
		},
	}
}
