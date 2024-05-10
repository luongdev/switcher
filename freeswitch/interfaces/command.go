package interfaces

type Command interface {
	Raw() string
}

type Filter interface {
}

type CommandOutput interface {
	IsOk() bool

	GetReply() string
}
