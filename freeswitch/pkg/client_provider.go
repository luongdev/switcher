package pkg

import (
	"github.com/luongdev/switcher/freeswitch/internal"
	"github.com/luongdev/switcher/freeswitch/types"
)

func NewClientProvider(store types.ClientStore) types.ClientProvider {
	return internal.NewClientProvider(store)
}
