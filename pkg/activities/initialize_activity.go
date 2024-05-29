package activities

import (
	"github.com/luongdev/switcher/internal/activities"
)

type InitializeActivityInput = activities.InitializeActivityInput
type InitializeActivityOutput = activities.InitializeActivityOutput

func NewInitializeActivity() *activities.InitializeActivity {
	return activities.NewInitializeActivity()
}
