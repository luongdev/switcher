package internal

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/percipia/eslgo"
	"log"
)

type ServerImpl struct {
	//cfg freeswitch.OutboundConfig

	host string
	port uint16
	ctx  context.Context

	sessionStarted interfaces.OnSessionFunc
	sessionEnded   interfaces.SessionEndedFunc
}

func (s *ServerImpl) OnSessionEnded(fn interfaces.SessionEndedFunc) {
	if fn != nil {
		s.sessionEnded = fn
	}
}

func (s *ServerImpl) OnSessionStarted(fn interfaces.OnSessionFunc) {
	if fn != nil {
		s.sessionStarted = fn
	}
}

func (s *ServerImpl) Start() error {
	hp := fmt.Sprintf("%v:%v", s.host, s.port)
	opts := eslgo.DefaultOutboundOptions

	opts.Context = s.ctx
	err := opts.ListenAndServe(hp, func(ctx context.Context, conn *eslgo.Conn, raw *eslgo.RawResponse) {
		if s.sessionStarted != nil {
			c := NewClient(conn, ctx)
			ss := NewSession(c, raw)

			s.sessionStarted(ctx, ss)
		}

		select {
		case <-ctx.Done():
			if s.sessionEnded != nil {
				s.sessionEnded()
			}
		}
	})

	return err
}

func NewServer(host string, port uint16, ctx context.Context) interfaces.Server {
	if port == 0 {
		port = 65022
	}

	if len(host) == 0 {
		host = "0.0.0.0"
	}

	s := &ServerImpl{host: host, port: port, ctx: ctx}
	s.sessionEnded = func() {
		log.Printf("Session closed")
	}

	return s
}

var _ interfaces.Server = (*ServerImpl)(nil)
