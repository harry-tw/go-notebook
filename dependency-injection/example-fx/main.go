package main

import (
	"context"
	"go-notebook/dependency-injection/config"
	"go-notebook/dependency-injection/repo"
	"go-notebook/dependency-injection/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log"
	"time"
)

func main() {
	app := fx.New(
		fx.Provide(
			newApplication,
			service.NewUserService,
			service.NewOrderService,
			service.NewCacheService,
			repo.NewUserRepo,
			repo.NewOrderRepo,
			provideUserDao,
			provideOrderDao,
			provideCacheLRU,
			config.NewConfig,
		),
		fx.Invoke(invoke),
		fx.WithLogger(
			func() fxevent.Logger {
				return fxevent.NopLogger
			},
		),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

// invoke is used to trigger fx to construct Application.
func invoke(app *Application) {
}

type Application struct {
	userService  *service.UserService
	orderService *service.OrderService
	cacheService *service.CacheService
}

func (a *Application) Serve() {
	log.Println("serve")
}

func (a *Application) Stop() {
	log.Println("stop")
}

func newApplication(lc fx.Lifecycle, userService *service.UserService, orderService *service.OrderService, cacheService *service.CacheService) (*Application, error) {
	app := &Application{
		userService:  userService,
		orderService: orderService,
		cacheService: cacheService,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("OnStart hook")
			app.Serve()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("OnStop hook")
			app.Stop()
			return nil
		},
	})
	return app, nil
}
