package activities

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal/activities"
)

type PlayActivityInput = activities.PlayActivityInput
type PlayActivityOutput = activities.PlayActivityOutput

func NewPlayActivity(provider types.ClientProvider) *activities.PlayActivity {
	return activities.NewPlayActivity(provider)
}
