package pkg

import (
	"github.com/luongdev/switcher/freeswitch/internal"
	"github.com/luongdev/switcher/freeswitch/types"
)

func NewClientStore(m map[string]types.Client) types.ClientStore {
	return internal.NewClientStore(m)
}
