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
