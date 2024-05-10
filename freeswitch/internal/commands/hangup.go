package commands

import (
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/percipia/eslgo/command/call"
)

type HangupCommand struct {
	call.Execute
}

func (a *HangupCommand) Raw() string {
	return a.Execute.BuildMessage()
}

func NewHangupCommand(uid string, cause string) *HangupCommand {
	return &HangupCommand{Execute: call.Execute{UUID: uid, AppName: "hangup", AppArgs: cause}}
}

var _ interfaces.Command = (*HangupCommand)(nil)
