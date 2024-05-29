package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/freeswitch/utils"
	"github.com/percipia/eslgo/command/call"
)

type LuaCommand struct {
	UId

	path string
}

func (c *LuaCommand) Validate() error {
	if err := c.UId.Validate(); err != nil {
		return err
	}

	if !utils.IsPathValid(c.path, ".lua") {
		return fmt.Errorf("path must be a .lua file. Got %v", c.path)
	}

	return nil
}

func (c *LuaCommand) Raw() (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: c.uid, AppName: "lua", AppArgs: fmt.Sprintf("%v", c.path)}).BuildMessage(), nil
}

func NewLuaCommand(uid, fileName string) *LuaCommand {
	return &LuaCommand{UId: UId{uid: uid, allowMissing: true}, path: fileName}
}

var _ types.Command = (*LuaCommand)(nil)
