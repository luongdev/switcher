package internal

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo"
)

type CommandImpl struct {
	cmd string
}

func (c *CommandImpl) BuildMessage() string {
	return c.cmd
}

func NewCommand(cmd string) *CommandImpl {
	return &CommandImpl{cmd: cmd}
}

type CommandOutputImpl struct {
	*eslgo.RawResponse
}

func (c *CommandOutputImpl) GetReply() string {
	return c.RawResponse.GetReply()
}

func (c *CommandOutputImpl) IsOk() bool {
	return c.RawResponse.IsOk()
}

func NewCommandOutput(raw *eslgo.RawResponse) *CommandOutputImpl {
	return &CommandOutputImpl{RawResponse: raw}
}

var _ types.CommandOutput = (*CommandOutputImpl)(nil)
