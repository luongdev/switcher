package activities

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal/activities"
)

type HangupActivityInput = activities.HangupActivityInput
type HangupActivityOutput = activities.HangupActivityOutput

func NewHangupActivity(provider types.ClientProvider) *activities.HangupActivity {
	return activities.NewHangupActivity(provider)
}
