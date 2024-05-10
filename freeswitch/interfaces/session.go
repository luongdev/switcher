package interfaces

import "context"

type Session interface {
	Event

	GetId() string

	Events(ctx context.Context, filters ...Filter)

	Answer(ctx context.Context) error
}
