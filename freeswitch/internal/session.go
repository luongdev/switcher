package internal

import (
	"context"
	"github.com/luongdev/switcher/freeswitch/internal/commands"
	"github.com/luongdev/switcher/freeswitch/types"
)

type SessionImpl struct {
	types.Event

	client types.Client
}

func (s *SessionImpl) Hangup(ctx context.Context, cause string) error {
	_, err := s.client.Exec(ctx, commands.NewHangupCommand(s.GetId(), cause))

	return err
}

func (s *SessionImpl) Exec(ctx context.Context, cmd types.Command) (types.CommandOutput, error) {
	return s.client.Exec(ctx, cmd)
}

func (s *SessionImpl) Answer(ctx context.Context) error {
	_, err := s.client.Exec(ctx, commands.NewAnswerCommand(s.GetId()))

	return err
}

func (s *SessionImpl) Events(ctx context.Context, filters ...types.Filter) {
}

func (s *SessionImpl) GetId() string {
	return s.Event.GetHeader("Unique-ID")
}

func NewSession(client types.Client, e types.Event) types.Session {
	return &SessionImpl{client: client, Event: e}
}

var _ types.Session = (*SessionImpl)(nil)
