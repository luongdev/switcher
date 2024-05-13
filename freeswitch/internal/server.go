package internal

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo"
	"log"
)

type ServerImpl struct {
	host           string
	port           uint16
	ctx            context.Context
	store          types.ClientStore
	sessionStarted types.OnSessionFunc
	sessionEnded   types.SessionEndedFunc
}

func (s *ServerImpl) SetStore(store types.ClientStore) {
	if store != nil {
		s.store = store
	}
}

func (s *ServerImpl) OnSessionEnded(fn types.SessionEndedFunc) {
	if fn != nil {
		s.sessionEnded = fn
	}
}

func (s *ServerImpl) OnSessionStarted(fn types.OnSessionFunc) {
	if fn != nil {
		s.sessionStarted = fn
	}
}

func (s *ServerImpl) Start() error {
	hp := fmt.Sprintf("%v:%v", s.host, s.port)
	opts := eslgo.DefaultOutboundOptions

	opts.Context = s.ctx
	err := opts.ListenAndServe(hp, func(ctx context.Context, conn *eslgo.Conn, raw *eslgo.RawResponse) {
		uid := raw.GetHeader("Unique-ID")
		sCtx := context.WithValue(ctx, "uid", uid)
		c := NewClient(conn, sCtx)

		if s.store != nil {
			s.store.Set(uid, c)
		}

		if s.sessionStarted != nil {
			ss := NewSession(c, raw)
			s.sessionStarted(sCtx, ss)
		}

		select {
		case <-ctx.Done():
			s.store.Del(uid)
			if s.sessionEnded != nil {
				s.sessionEnded()
			}
		}
	})

	return err
}

func NewServer(host string, port uint16) types.Server {
	if port == 0 {
		port = 65022
	}

	if len(host) == 0 {
		host = "0.0.0.0"
	}

	s := &ServerImpl{host: host, port: port, ctx: context.Background()}
	s.sessionEnded = func() {
		log.Printf("Session closed")
	}

	return s
}

var _ types.Server = (*ServerImpl)(nil)
