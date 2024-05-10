package pkg

import (
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/luongdev/switcher/freeswitch/internal/commands"
)

func AnswerCommand(uid string) interfaces.Command {
	return commands.NewAnswerCommand(uid)
}
