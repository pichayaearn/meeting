package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/pichayaearn/meeting/cmd/api/config"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("api starting ...")
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	log := logrus.StandardLogger()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Start Login API server.
	s := newServer(cfg)

	log.Info("starting server...")

	go func() {
		log.Error(s.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)))
	}()

	<-quit

	log.Info("receive signal, stopping Login API...")

	ctx, cancel := context.WithCancel(context.Background())
	// Shutdown Server
	if err := s.Shutdown(ctx); err != nil {
		log.Error(err)
	}

	cancel()
}
