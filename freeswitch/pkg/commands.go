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

func OriginateCommand(background bool, aleg, bleg *types.Leg, vars map[string]interface{}) types.Command {
	return commands.NewOriginateCommand(background, aleg, bleg, vars)
}

func BridgeCommand(uid string, otherLeg interface{}) types.Command {
	return commands.NewBridgeCommand(uid, otherLeg)
}

func LuaCommand(uid, fileName string) types.Command {
	return commands.NewLuaCommand(uid, fileName)
}

func ReloadCommand(reType string) types.Command {
	return commands.NewReloadCommand(reType)
}

func LoadCommand(module string, unload bool) types.Command {
	return commands.NewLoadCommand(module, unload)
}

func EchoCommand(uid string) types.Command {
	return commands.NewEchoCommand(uid)
}
