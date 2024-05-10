package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/percipia/eslgo/command/call"
)

type SetCommand struct {
	call.Execute
}

func (a *SetCommand) Raw() string {
	return a.Execute.BuildMessage()
}

func NewSetCommand(uid string, vars map[string]interface{}) *SetCommand {
	cmd := "set"
	args := ""

	del := ";"
	multi := len(vars) > 1
	if multi {
		cmd = "multiset"
		args = fmt.Sprintf("^^")
	}

	for k, v := range vars {
		arg := fmt.Sprintf("%s=%v", k, v)
		if multi {
			args += fmt.Sprintf("%s%s", del, arg)
			continue
		}
		args = arg
	}

	return &SetCommand{Execute: call.Execute{UUID: uid, AppName: cmd, AppArgs: args}}
}

var _ interfaces.Command = (*SetCommand)(nil)
