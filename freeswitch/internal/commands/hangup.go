package commands

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command"
	"github.com/percipia/eslgo/command/call"
)

type HangupCommand struct {
	UId

	cause string
}

func (a *HangupCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	var cmd command.Command
	if a.uid != "" {
		cmd = &call.Execute{UUID: a.uid, AppName: "hangup", AppArgs: a.cause}
	} else {
		cmd = &command.API{Command: "hupall", Arguments: a.cause}
	}

	return cmd.BuildMessage(), nil
}

func NewHangupCommand(uid string, cause string) *HangupCommand {
	return &HangupCommand{UId: UId{uid: uid, allowMissing: true}, cause: cause}
}

var _ types.Command = (*HangupCommand)(nil)
