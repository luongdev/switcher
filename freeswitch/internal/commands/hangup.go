package commands

import (
	"github.com/luongdev/switcher/freeswitch/types"
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

	appName := "hupall"
	if a.uid != "" {
		appName = "hangup"
	}

	return (&call.Execute{UUID: a.uid, AppName: appName, AppArgs: a.cause}).BuildMessage(), nil
}

func NewHangupCommand(uid string, cause string) *HangupCommand {
	return &HangupCommand{UId: UId{uid: uid, allowMissing: true}, cause: cause}
}

var _ types.Command = (*HangupCommand)(nil)
