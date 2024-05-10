package types

import "context"

type Session interface {
	Event

	GetId() string

	Events(ctx context.Context, filters ...Filter)

	Answer(ctx context.Context) error

	Hangup(ctx context.Context, cause string) error

	Exec(ctx context.Context, cmd Command) (CommandOutput, error)
}
