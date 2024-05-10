package pkg

import (
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/luongdev/switcher/freeswitch/internal"
)

func NewClientStore(m map[string]interfaces.Client) interfaces.ClientStore {
	return internal.NewClientStore(m)
}
