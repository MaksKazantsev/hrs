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
	l := log.NewLogger(cfg.Env)
	l.Info("starting app", slog.Int("port", cfg.Port), slog.Any("config:", cfg))

	defer func() {
		if err := recover(); err != nil {
			l.Error("recovery failed ❌", slog.Any("error", err))
		}
	}()

	// New db
	db := postgres.MustConnect(cfg.Database.GetDSN())
	defer func() {
		_ = db.Close()
	}()

	// Init repo
	repo := postgres.NewRepository(db)
	l.Info("successfully connected to db ✔")

	// Init service
	srvc := service.NewService(repo)

	// New GRPC server
	serv := server.NewServer(l, srvc)
	l.Info("all layers successfully set up ✔")

	// Run server
	shutdown(func() {
		l.Info("server is running ✔")
		run(cfg.Port, serv)
	})
	l.Info("server stopped ✔")
}

func run(port int, s *grpc.Server) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic("failed to listen tcp: " + err.Error())
	}

	if err = s.Serve(l); err != nil {
		panic("failed to start server: " + err.Error())
	}
}

func shutdown(fn func()) {
	chStop := make(chan os.Signal, 1)
	signal.Notify(chStop, syscall.SIGINT, syscall.SIGTERM)
	go fn()
	<-chStop
}
