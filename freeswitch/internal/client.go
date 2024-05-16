package internal

import (
	"context"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo"
)

type ClientImpl struct {
	conn      *eslgo.Conn
	sessionId string
	ctx       context.Context
}

func (c *ClientImpl) GetSessionId() string {
	return c.sessionId
}

func (c *ClientImpl) Disconnect() {
	if c.conn != nil {
		c.conn.ExitAndClose()
	}
}

func (c *ClientImpl) Exec(ctx context.Context, cmd types.Command) (types.CommandOutput, error) {
	in, err := cmd.Raw()
	if err != nil {
		return nil, err
	}

	out, err := c.conn.SendCommand(ctx, NewCommand(in))
	if err != nil {
		return nil, err
	}

	return NewCommandOutput(out), nil
}

func (c *ClientImpl) Events(ctx context.Context) error {
	return nil
}

func NewClient(c *eslgo.Conn, ctx context.Context) *ClientImpl {
	//c.OriginateCall()
	return &ClientImpl{conn: c, ctx: ctx}
}
func NewSessionClient(c *eslgo.Conn, sessionId string, ctx context.Context) *ClientImpl {
	return &ClientImpl{conn: c, ctx: ctx, sessionId: sessionId}
}

var _ types.Client = (*ClientImpl)(nil)
