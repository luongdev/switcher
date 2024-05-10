package interfaces

import "context"

type Server interface {
	Start() error

	SetStore(store ClientStore)

	OnSessionStarted(fn OnSessionFunc)
	OnSessionEnded(fn SessionEndedFunc)
}

type OnSessionFunc func(ctx context.Context, session Session)
type SessionEndedFunc func()
