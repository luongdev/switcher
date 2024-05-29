package activities

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal/activities"
)

type SetActivityInput = activities.SetActivityInput
type SetActivityOutput = activities.SetActivityOutput

func NewSetActivity(provider types.ClientProvider) *activities.SetActivity {
	return activities.NewSetActivity(provider)
}
