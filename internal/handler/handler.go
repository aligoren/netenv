package handler

import (
	"bufio"
	"io"
	"net"
	"strings"

	"github.com/aligoren/netenv/config"
	"github.com/aligoren/netenv/internal/command"
	"github.com/aligoren/netenv/pkg/logger"
)

type Request struct {
	Con net.Conn
}

func (r Request) HandleRequest() {

	conf, err := config.GetConfig()
	if err != nil {
		logger.Get().Error(err)
	}

	defer r.Con.Close()

	reader := bufio.NewReader(r.Con)

	for {
		request, err := reader.ReadString('\n')

		switch err {
		case nil:
			request = strings.TrimSpace(request)
		case io.EOF:
			logger.Get().Info("client closed the connection by terminating the process")
			return
		default:
			logger.Get().Error("error: ", err)
			return
		}

		c := command.Command{
			CommandText: request,
			Config:      conf,
			Con:         r.Con,
		}

		response := c.HandleCommand()
		if _, err := r.Con.Write([]byte(response)); err != nil {
			logger.Get().Error("error: ", err)
		}
	}
}
