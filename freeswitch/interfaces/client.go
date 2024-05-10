package interfaces

import "context"

type Client interface {
	Disconnect()

	Exec(ctx context.Context, cmd Command) (CommandOutput, error)
	Events(ctx context.Context) error
}

type ClientDisconnectFunc func(ctx context.Context)
