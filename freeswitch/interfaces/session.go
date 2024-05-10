package interfaces

type Session interface {
	Event

	GetId() string
}
