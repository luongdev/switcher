package pkg

import (
	"github.com/luongdev/switcher/freeswitch/internal/commands"
	"github.com/luongdev/switcher/freeswitch/types"
)

func AnswerCommand(uid string) types.Command {
	return commands.NewAnswerCommand(uid)
}

func HangupCommand(uid, cause string) types.Command {
	return commands.NewHangupCommand(uid, cause)
}

func SetCommand(uid string, vars map[string]interface{}) types.Command {
	return commands.NewSetCommand(uid, vars)
}

func NewOriginateCommand(background bool, aleg, bleg *types.Leg, vars map[string]interface{}) types.Command {
	return commands.NewOriginateCommand(background, aleg, bleg, vars)
}

func NewBridgeCommand(uid string, otherLeg *types.Leg) types.Command {
	return commands.NewBridgeCommand(uid, otherLeg)
}

func NewLuaCommand(uid, fileName string) types.Command {
	return commands.NewLuaCommand(uid, fileName)
}

func NewReloadCommand(reType string) types.Command {
	return commands.NewReloadCommand(reType)
}

func NewLoadCommand(module string, unload bool) types.Command {
	return commands.NewLoadCommand(module, unload)
}

func NewEchoCommand(uid string) types.Command {
	return commands.NewEchoCommand(uid)
}
