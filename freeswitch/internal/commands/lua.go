package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command/call"
	"path/filepath"
)

type LuaCommand struct {
	UId

	fileName string
}

func (c *LuaCommand) Validate() error {
	if err := c.UId.Validate(); err != nil {
		return err
	}

	if c.fileName == "" {
		return fmt.Errorf("invalid file name")
	}

	if filepath.Ext(c.fileName) != ".lua" {
		return fmt.Errorf("invalid file extension")
	}

	return nil
}

func (c *LuaCommand) Raw() (string, error) {
	if err := c.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: c.uid, AppName: "lua", AppArgs: fmt.Sprintf("%v", c.fileName)}).BuildMessage(), nil
}

func NewLuaCommand(uid, fileName string) *LuaCommand {
	return &LuaCommand{UId: UId{uid: uid, allowMissing: true}, fileName: fileName}
}

var _ types.Command = (*LuaCommand)(nil)
