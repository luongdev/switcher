package commands

import (
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/percipia/eslgo/command/call"
)

type AnswerCommand struct {
	call.Execute
}

func (a *AnswerCommand) Raw() string {
	return a.Execute.BuildMessage()
}

func NewAnswerCommand(uid string) *AnswerCommand {
	return &AnswerCommand{Execute: call.Execute{UUID: uid, AppName: "answer"}}
}

var _ interfaces.Command = (*AnswerCommand)(nil)
