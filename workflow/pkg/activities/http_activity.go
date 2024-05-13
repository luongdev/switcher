package activities

import (
	"github.com/luongdev/switcher/workflow/internal/activities"
	"github.com/luongdev/switcher/workflow/types"
)

type HttpActivityInput activities.HttpActivityInput

func HttpActivity() types.Activity {
	return &activities.HttpActivity{}
}
