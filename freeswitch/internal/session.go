package internal

import "github.com/luongdev/switcher/freeswitch/interfaces"

type SessionImpl struct {
	interfaces.Event

	client interfaces.Client
}

func (s *SessionImpl) GetId() string {
	return s.Event.GetHeader("Unique-ID")
}

func NewSession(client interfaces.Client, e interfaces.Event) interfaces.Session {
	return &SessionImpl{client: client, Event: e}
}

var _ interfaces.Session = (*SessionImpl)(nil)
