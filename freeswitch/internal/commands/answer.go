package commands

import (
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command/call"
)

type AnswerCommand struct {
	UIdCommand
}

func (a *AnswerCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: a.uid, AppName: "answer"}).BuildMessage(), nil
}

func NewAnswerCommand(uid string) *AnswerCommand {
	return &AnswerCommand{UIdCommand: UIdCommand{uid: uid}}
}

var _ types.Command = (*AnswerCommand)(nil)
