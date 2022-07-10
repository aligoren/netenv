package main

import (
	"github.com/aligoren/netenv/internal/server"
	"github.com/aligoren/netenv/pkg/logger"
)

func main() {

	server := server.New(":8080")

	err := server.Start()
	if err != nil {
		logger.Get().Error(err)
	}
}
