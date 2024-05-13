package internal

import "github.com/luongdev/switcher/freeswitch/types"

type ClientStoreImpl struct {
	m map[string]types.Client
}

func keyOrDefault(k string) string {
	if len(k) == 0 {
		k = types.DefaultClient
	}
	return k
}

func (s *ClientStoreImpl) Set(k string, v types.Client) {
	k = keyOrDefault(k)
	if v != nil {
		s.m[k] = v
	}
}

func (s *ClientStoreImpl) Del(k string) {
	k = keyOrDefault(k)
	if k == types.DefaultClient {
		return
	}

	if _, ok := s.m[k]; ok {
		delete(s.m, k)
	}
}

func (s *ClientStoreImpl) Get(k string) (types.Client, bool) {
	k = keyOrDefault(k)
	v, ok := s.m[k]

	return v, ok
}

func NewClientStore(m map[string]types.Client) types.ClientStore {
	if m == nil {
		m = make(map[string]types.Client)
	}
	return &ClientStoreImpl{m: m}
}

var _ types.ClientStore = (*ClientStoreImpl)(nil)
