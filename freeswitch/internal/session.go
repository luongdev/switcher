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
