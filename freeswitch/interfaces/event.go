package interfaces

type Event interface {
	HasHeader(name string) bool
	GetHeader(name string) string
	GetVariable(name string) string
}
