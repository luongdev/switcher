package internal

import (
	freeswitchtypes "github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/types"
)

type SessionActivity struct {
	Provider freeswitchtypes.ClientProvider
}

func (a *SessionActivity) GetClient(input types.SessionInput) freeswitchtypes.Client {
	return a.Provider.Get(input.GetSessionId())
}

func NewFreeswitchActivity(provider freeswitchtypes.ClientProvider) *SessionActivity {
	return &SessionActivity{Provider: provider}
}
