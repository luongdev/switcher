package pkg

import (
	"github.com/luongdev/switcher/freeswitch"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/luongdev/switcher/freeswitch/internal"
)

func NewClient(c freeswitch.InboundConfig) interfaces.Client {
	return internal.NewClient(c)
}
