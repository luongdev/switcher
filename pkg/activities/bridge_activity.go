package activities

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal/activities"
)

type BridgeActivityInput = activities.BridgeActivityInput
type BridgeActivityOutput = activities.BridgeActivityOutput

func NewBridgeActivity(provider types.ClientProvider) *activities.BridgeActivity {
	return activities.NewBridgeActivity(provider)
}
