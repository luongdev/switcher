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

func (c *ClientImpl) Api(ctx context.Context, api *interfaces.API) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientImpl) Execute(ctx context.Context, cmd *interfaces.Command) {
	//TODO implement me
	panic("implement me")
}

func (c *ClientImpl) Events(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func NewClient(c *eslgo.Conn, ctx context.Context) *ClientImpl {
	return &ClientImpl{conn: c, ctx: ctx}
}

var _ interfaces.Client = (*ClientImpl)(nil)
