package internal

import (
	"context"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/luongdev/switcher/freeswitch/internal/commands"
)

type SessionImpl struct {
	interfaces.Event

	client interfaces.Client
}

func (s *SessionImpl) Hangup(ctx context.Context, cause string) error {
	_, err := s.client.Exec(ctx, commands.NewHangupCommand(s.GetId(), cause))

	return err
}

func (s *SessionImpl) Exec(ctx context.Context, cmd interfaces.Command) (interfaces.CommandOutput, error) {
	return s.client.Exec(ctx, cmd)
}

func (s *SessionImpl) Answer(ctx context.Context) error {
	_, err := s.client.Exec(ctx, commands.NewAnswerCommand(s.GetId()))

	return err
}

func (s *SessionImpl) Events(ctx context.Context, filters ...interfaces.Filter) {
}

func (s *SessionImpl) GetId() string {
	return s.Event.GetHeader("Unique-ID")
}

func NewSession(client interfaces.Client, e interfaces.Event) interfaces.Session {
	return &SessionImpl{client: client, Event: e}
}

var _ interfaces.Session = (*SessionImpl)(nil)
