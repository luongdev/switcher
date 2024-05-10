package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo"
	"github.com/percipia/eslgo/command"
)

type OriginateCommand struct {
	aleg, bleg *types.Leg
	background bool
	vars       map[string]interface{}
}

func (a *OriginateCommand) Validate() error {
	if a.aleg == nil || !a.aleg.Valid() {
		return fmt.Errorf("aleg is not valid")
	}

	if a.bleg == nil || !a.bleg.Valid() {
		return fmt.Errorf("bleg is not valid")
	}

	return nil
}

func (a *OriginateCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	vars := make(map[string]string)
	for k, v := range a.vars {
		if k == "origination_uuid" {
			continue
		}
		vars[k] = fmt.Sprintf("%v", v)
	}

	api := &command.API{
		Command:    "originate",
		Arguments:  fmt.Sprintf("%s%s %s", eslgo.BuildVars("{%s}", vars), a.aleg.DialString(), a.bleg.DialString()),
		Background: a.background,
	}

	return api.BuildMessage(), nil
}

func NewOriginateCommand(background bool, aleg, bleg *types.Leg, vars map[string]interface{}) *OriginateCommand {
	if vars == nil {
		vars = make(map[string]interface{})
	}

	return &OriginateCommand{background: background, aleg: aleg, bleg: bleg, vars: vars}
}

var _ types.Command = (*OriginateCommand)(nil)
