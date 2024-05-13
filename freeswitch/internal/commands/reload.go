package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command"
	"strings"
)

type ReloadCommand struct {
	reload string
}

func (a *ReloadCommand) Validate() error {
	if a.reload == "" {
		return fmt.Errorf("module type is required")
	}

	if a.reload != "xml" && a.reload != "acl" && strings.Index(a.reload, "mod_") == -1 {
		return fmt.Errorf("invalid module type. Must be 'xml' or 'acl' or module name")
	}

	return nil
}

func (a *ReloadCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	cmd := "module"
	args := ""
	if a.reload != "xml" {
		cmd = "reloadxml"
	} else if a.reload != "acl" {
		cmd = "reloadacl"
	} else {
		cmd = "module"
		args = a.reload
	}

	return (&command.API{Command: cmd, Arguments: args}).BuildMessage(), nil
}

func NewReloadCommand(reload string) *ReloadCommand {
	return &ReloadCommand{reload: reload}
}

var _ types.Command = (*ReloadCommand)(nil)
