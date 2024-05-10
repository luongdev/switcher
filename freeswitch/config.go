package freeswitch

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/luongdev/switcher/freeswitch/internal"
	"github.com/percipia/eslgo"
	"time"
)

type InboundConfig struct {
	Host           string        `yaml:"host"`
	Port           uint16        `yaml:"port"`
	Password       string        `yaml:"password"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
}

type OutboundConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type Config struct {
	Inbound  InboundConfig  `yaml:"inbound"`
	Outbound OutboundConfig `yaml:"outbound"`
}

func (c *InboundConfig) Build() (interfaces.Client, error) {
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

	opts := eslgo.DefaultInboundOptions
	opts.Password = c.Password
	opts.AuthTimeout = c.ConnectTimeout

	ctx := context.WithValue(context.Background(), "opts", opts)
	opts.Options.Context = ctx

	hp := fmt.Sprintf("%v:%v", c.Host, c.Port)
	conn, err := opts.Dial(hp)
	if err != nil {
		return nil, err
	}

	return internal.NewClient(conn, ctx), nil
}

func (c *OutboundConfig) Build() interfaces.Server {
	return internal.NewServer(c.Host, c.Port, context.Background())
}
