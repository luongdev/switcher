package internal

import "github.com/luongdev/switcher/freeswitch/interfaces"

type ClientStoreImpl struct {
	m map[string]interfaces.Client
}

func keyOrDefault(k string) string {
	if len(k) == 0 {
		k = interfaces.DefaultClient
	}
	return k
}

func (s *ClientStoreImpl) Set(k string, v interfaces.Client) {
	k = keyOrDefault(k)
	if v != nil {
		s.m[k] = v
	}
}

func (s *ClientStoreImpl) Del(k string) {
	k = keyOrDefault(k)
	if k == interfaces.DefaultClient {
		return
	}

	if _, ok := s.m[k]; ok {
		delete(s.m, k)
	}
}

func (s *ClientStoreImpl) Get(k string) (interfaces.Client, bool) {
	k = keyOrDefault(k)
	v, ok := s.m[k]

	return v, ok
}

func NewClientStore(m map[string]interfaces.Client) interfaces.ClientStore {
	if m == nil {
		m = make(map[string]interfaces.Client)
	}
	return &ClientStoreImpl{m: m}
}

var _ interfaces.ClientStore = (*ClientStoreImpl)(nil)
