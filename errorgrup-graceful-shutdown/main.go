package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-notebook/errorgrup-graceful-shutdown/model"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// catch signals
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
		<-sigs
		log.Println("receive termination signal")
		cancel()
	}()

	// setup errgroup
	errGrp, errGrpCtx := errgroup.WithContext(ctx)
	// run server
	errGrp.Go(func() error {
		log.Println("run server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	// gracefully shutdown server based on errgroup context
	errGrp.Go(func() error {
		select {
		case <-errGrpCtx.Done():
			if err := srv.Shutdown(context.Background()); err != nil {
				return err
			}
			log.Println("shutdown server")
		}
		return nil
	})
	// run pprof
	pprof := http.Server{Addr: "localhost:6060"}
	errGrp.Go(func() error {
		log.Println("run pprof")
		if err := pprof.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	// gracefully shutdown pprof based on errgroup context
	errGrp.Go(func() error {
		select {
		case <-errGrpCtx.Done():
			if err := pprof.Shutdown(context.Background()); err != nil {
				return err
			}
			log.Println("shutdown pprof")
		}
		return nil
	})
	if err := errGrp.Wait(); err != nil {
		log.Panic(err)
	}

	log.Println("quit")
}

func NewServer() http.Server {
	engine := gin.Default()
	setupRouter(engine)
	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: engine,
	}
	return srv
}

func setupRouter(engine *gin.Engine) {
	engine.POST("/hello", func(c *gin.Context) {
		var req model.HelloReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, model.ErrRsp{Err: "invalid request"})
			return
		}
		c.JSON(http.StatusOK, model.HelloRsp{Msg: fmt.Sprintf("hello %s", req.Name)})
	})
}
