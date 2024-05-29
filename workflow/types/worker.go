package types

type Worker interface {
	Start() error

	Stop()
}
