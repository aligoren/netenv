package server

import (
	"net"

	"github.com/aligoren/netenv/internal/handler"
	"github.com/aligoren/netenv/pkg/logger"
)

type Server struct {
	addr string
}

func New(addr string) Server {
	return Server{
		addr,
	}
}

func (s Server) Start() error {

	r := &handler.Request{}

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			logger.Get().Error(err)
			continue
		}

		r.Con = con

		go r.HandleRequest()
	}

}
