package server

import (
	"fmt"
	"net/http"

	"github.com/leagueify/communication/handler"
	"github.com/leagueify/communication/internal/config"
	"github.com/leagueify/communication/internal/middleware"
)

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Start() {
	router := http.NewServeMux()
	communicationRouter := handler.CommunicationRouter()

	router.Handle(
		"/communication/", http.StripPrefix(
			"/communication", communicationRouter,
		),
	)

	// define server config
	communicationServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler:      middleware.Logging(router),
		IdleTimeout:  s.cfg.Server.IdleTimeout,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	// show server banner
	showStartBanner()

	// start server
	if err := communicationServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
