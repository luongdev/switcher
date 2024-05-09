package internal

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/freeswitch"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/percipia/eslgo"
	"time"
)

type ClientImpl struct {
	conn *eslgo.Conn

	cfg          freeswitch.InboundConfig
	onDisconnect interfaces.ClientDisconnectFunc
}

func (c *ClientImpl) OnDisconnect(fn interfaces.ClientDisconnectFunc) {
	if fn != nil {
		c.onDisconnect = fn
	}
}

func (c *ClientImpl) Connect(ctx context.Context) error {
	opts := eslgo.DefaultInboundOptions
	opts.Options.Context = ctx
	opts.Password = c.cfg.Password
	opts.AuthTimeout = c.cfg.ConnectTimeout
	opts.OnDisconnect = func() {
		if c.onDisconnect != nil {
			c.onDisconnect(ctx)
		}
	}

	hp := fmt.Sprintf("%v:%v", c.cfg.Host, c.cfg.Port)
	conn, err := opts.Dial(hp)
	if err != nil {
		return err
	}

	c.conn = conn

	return nil
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

func NewClient(c freeswitch.InboundConfig) *ClientImpl {
	if len(c.Host) == 0 {
		c.Host = "127.0.0.1"
	}

	if c.Port == 0 {
		c.Port = 65021
	}

	if len(c.Password) == 0 {
		c.Password = "Simplefs!!"
	}

	if c.ConnectTimeout < time.Second*1 {
		c.ConnectTimeout = time.Second
	}

	client := &ClientImpl{
		cfg:          c,
		onDisconnect: func(ctx context.Context) {},
	}

	return client
}

var _ interfaces.Client = (*ClientImpl)(nil)
