package app

import (
	"fmt"
	"github.com/alserov/hrs/auth/internal/config"
	"github.com/alserov/hrs/auth/internal/db/postgres"
	"github.com/alserov/hrs/auth/internal/log"
	"github.com/alserov/hrs/auth/internal/server"
	"github.com/alserov/hrs/auth/internal/service"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	// Init logger
	l := log.MustSetup(cfg.Env)
	l.Info("starting app", slog.Int("port", cfg.Port))

	// New db
	db := postgres.MustConnect(cfg.Database.GetDSN())
	defer db.Close()

	// Init repo

	repo := postgres.NewRepository(db)

	// Init service
	srvc := service.NewService(repo)

	// New GRPC server
	serv := grpc.NewServer()
	server.RegisterGRPCServer(serv, srvc)

	// Run server
	l.Info("app is running")
	run(serv, cfg.Port)
	l.Info("app was stopped")
}

func run(s *grpc.Server, port int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = s.Serve(l); err != nil {
			panic("failed to serve: " + err.Error())
		}
	}()

	<-chStop
}
