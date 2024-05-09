package interfaces

import "context"

type Client interface {
	Connect(ctx context.Context) error
	Disconnect()
	OnDisconnect(fn ClientDisconnectFunc)

	Api(ctx context.Context, api *API)
	Execute(ctx context.Context, cmd *Command)
	Events(ctx context.Context)
}

type ClientDisconnectFunc func(ctx context.Context)
