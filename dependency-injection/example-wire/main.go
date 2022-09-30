package main

import (
	"go-notebook/dependency-injection/service"
	"log"
)

func main() {
	app, err := InitializeApplication()
	if err != nil {
		log.Fatal(err)
	}
	app.Serve()
}

type Application struct {
	cacheService *service.CacheService
	userService  *service.UserService
	orderService *service.OrderService
}

func (app *Application) Serve() {
	log.Println("serve")
}

func newApplication(cacheService *service.CacheService, userService *service.UserService, orderService *service.OrderService) *Application {
	return &Application{
		cacheService: cacheService,
		userService:  userService,
		orderService: orderService,
	}
}
