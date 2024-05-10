package interfaces

import "context"

type Server interface {
	Start() error

	OnSessionStarted(fn OnSessionFunc)
	OnSessionEnded(fn SessionEndedFunc)
}

type OnSessionFunc func(ctx context.Context, session Session)
type SessionEndedFunc func()
