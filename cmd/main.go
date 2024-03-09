package main

import (
	"context"
	"flag"
	"github.com/RostislavOrlov/krp_admin/internal/db"
	"github.com/RostislavOrlov/krp_admin/internal/handlers/user"
	"github.com/RostislavOrlov/krp_admin/internal/repositories"
	"github.com/RostislavOrlov/krp_admin/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var listen string

	flag.StringVar(&listen, "listen", ":8082", "server listen interface")

	flag.Parse()

	ctx := context.Background()
	//pgDB, err := db.NewPostgresPool("postgres://postgres:12345@localhost:5432/krp_auth")
	pgDB, err := db.NewPostgresPool("postgres://krp_auth:krp_auth@postgres_auth:5432/krp_auth")
	if err != nil {
		logrus.Panicf("unable get postgres pool: %v", err)
	}

	userRepo, err := repositories.NewUserRepository(pgDB)
	if err != nil {
		logrus.Panicf("unable build user repo: %v", err)
	}

	userService, err := services.NewUserService(userRepo)
	if err != nil {
		logrus.Panicf("unable build user service: %v", err)
	}

	g := gin.New()

	_, err = user.NewUserHandler(userService, g)
	if err != nil {
		logrus.Panicf("unable build user handler: %v", err)
	}

	doneC := make(chan error)

	go func() { doneC <- g.Run(listen) }()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGTERM)

	childCtx, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-signalChan
		logrus.Debugf("exiting with signal: %v", sig)
		cancel()
	}()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				doneC <- ctx.Err()
			}
		}
	}(childCtx)

	<-doneC
}
