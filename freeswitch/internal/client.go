package internal

import (
	"context"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/percipia/eslgo"
)

type ClientImpl struct {
	conn *eslgo.Conn

	ctx context.Context
}

func (c *ClientImpl) Disconnect() {
	if c.conn != nil {
		c.conn.ExitAndClose()
	}
}

func (c *ClientImpl) Exec(ctx context.Context, cmd interfaces.Command) (interfaces.CommandOutput, error) {
	raw, err := c.conn.SendCommand(ctx, NewCommand(cmd.Raw()))
	if err != nil {
		return nil, err
	}

	return NewCommandOutput(raw), nil
}

func (c *ClientImpl) Events(ctx context.Context) error {
	return nil
}

func NewClient(c *eslgo.Conn, ctx context.Context) *ClientImpl {
	return &ClientImpl{conn: c, ctx: ctx}
}

var _ interfaces.Client = (*ClientImpl)(nil)
