package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command"
	"strings"
)

type LoadCommand struct {
	module string
	unload bool
}

func (a *LoadCommand) Validate() error {
	if a.module == "" || strings.Index(a.module, "mod_") == -1 {
		return fmt.Errorf("module type is required")
	}

	return nil
}

func (a *LoadCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	cmd := "load"
	if a.unload {
		cmd = "unload"
	}

	return (&command.API{Command: cmd, Arguments: a.module}).BuildMessage(), nil
}

func NewLoadCommand(module string, unload bool) *LoadCommand {
	return &LoadCommand{module: module, unload: unload}
}

var _ types.Command = (*LoadCommand)(nil)
