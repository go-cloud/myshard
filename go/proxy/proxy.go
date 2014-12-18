package proxy

import (
	. "github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/server"
	"github.com/siddontang/go/log"
	"net"
)

// Server acts as a MySQL server, it listens and accepts MySQL client connections,
// parses the commands then sends them to the specified MySQL server.
type Server struct {
	l net.Listener

	quit     chan struct{}
	addr     string
	user     string
	password string
}

func NewServer(addr string, user string, password string) (*Server, error) {
	s := new(Server)

	s.addr = addr
	s.user = user
	s.password = password

	s.quit = make(chan struct{})

	var err error
	proto := GetNetProto(addr)
	if s.l, err = net.Listen(proto, addr); err != nil {
		return nil, err
	}

	go s.run()

	return s, nil
}

func (s *Server) run() {
	for {
		c, err := s.l.Accept()
		if err != nil {
			return
		}

		go s.onConn(c)
	}
}

func (s *Server) onConn(c net.Conn) {
	conn, err := server.NewConn(c, s.user, s.password, new(Handler))
	if err != nil {
		log.Error("new connection error %s", err.Error())
		c.Close()
		return
	}

	for {
		err = conn.HandleCommand()
		if err != nil {
			log.Error("handle command error %s", err.Error())
			return
		}
	}
}

func (s *Server) Close() {
	close(s.quit)

	if s.l != nil {
		s.l.Close()
	}
}
