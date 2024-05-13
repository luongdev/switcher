package types

type Command interface {
	Raw() (string, error)

	Validate() error
}

type Filter interface {
}

type CommandOutput interface {
	IsOk() bool

	GetReply() string
}
