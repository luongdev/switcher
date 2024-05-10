package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command/call"
)

type SetCommand struct {
	UId

	vars  map[string]interface{}
	multi bool
}

func (a *SetCommand) Validate() error {
	err := a.UId.Validate()
	if err != nil {
		return err
	}

	if len(a.vars) == 0 {
		return fmt.Errorf("at least variable is required")
	}

	return nil
}

func (a *SetCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	cmd := "set"
	args := ""

	del := ";"
	if a.multi {
		cmd = "multiset"
		args = fmt.Sprintf("^^")
	}

	for k, v := range a.vars {
		arg := fmt.Sprintf("%s=%v", k, v)
		if a.multi {
			args += fmt.Sprintf("%s%s", del, arg)
			continue
		}
		args = arg
	}

	return (&call.Execute{UUID: a.uid, AppName: cmd, AppArgs: args}).BuildMessage(), nil
}

func NewSetCommand(uid string, vars map[string]interface{}) *SetCommand {

	return &SetCommand{UId: UId{uid: uid}, vars: vars, multi: len(vars) > 1}
}

var _ types.Command = (*SetCommand)(nil)
