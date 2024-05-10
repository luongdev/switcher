package pkg

import (
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/luongdev/switcher/freeswitch/internal/commands"
)

func AnswerCommand(uid string) interfaces.Command {
	return commands.NewAnswerCommand(uid)
}

func HangupCommand(uid, cause string) interfaces.Command {
	return commands.NewHangupCommand(uid, cause)
}

func SetCommand(uid string, vars map[string]interface{}) interfaces.Command {
	return commands.NewSetCommand(uid, vars)
}
