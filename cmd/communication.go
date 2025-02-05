package main

import (
	"github.com/leagueify/communication/internal/config"
	"github.com/leagueify/communication/internal/integrations/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewServer(cfg).Start()
}
