package commands

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command/call"
)

type EchoCommand struct {
	UId
}

func (a *EchoCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: a.uid, AppName: "echo"}).BuildMessage(), nil
}

func NewEchoCommand(uid string) *EchoCommand {
	return &EchoCommand{UId: UId{uid: uid}}
}

var _ types.Command = (*EchoCommand)(nil)
