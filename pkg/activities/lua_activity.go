package activities

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal/activities"
)

type LuaActivityInput = activities.LuaActivityInput
type LuaActivityOutput = activities.LuaActivityOutput

func NewLuaActivity(provider types.ClientProvider) *activities.LuaActivity {
	return activities.NewLuaActivity(provider)
}
