package internal

import (
	freeswitchtypes "github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/types"
)

type SessionActivity struct {
	provider freeswitchtypes.ClientProvider
}

func (a *SessionActivity) GetClient(input types.SessionInput) (freeswitchtypes.Client, bool) {
	return a.provider.Get(input.GetSessionId())
}

func NewFreeswitchActivity(provider freeswitchtypes.ClientProvider) SessionActivity {
	return SessionActivity{provider: provider}
}
