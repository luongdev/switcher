package internal

import "github.com/luongdev/switcher/freeswitch/types"

type ClientProviderImpl struct {
	store types.ClientStore
}

func (c *ClientProviderImpl) Get(k string) (types.Client, bool) {
	return c.store.Get(k)
}

func NewClientProvider(store types.ClientStore) types.ClientProvider {
	return &ClientProviderImpl{store: store}
}

var _ types.ClientProvider = (*ClientProviderImpl)(nil)
