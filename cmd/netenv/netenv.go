package main

import (
	"github.com/aligoren/netenv/config"
	"github.com/aligoren/netenv/internal/server"
	"github.com/aligoren/netenv/pkg/logger"
)

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Get().Error(err)
	}

	server := server.New(cfg.Global.Addr)

	err = server.Start()
	if err != nil {
		logger.Get().Error(err)
	}
}
