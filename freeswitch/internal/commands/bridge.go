package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/percipia/eslgo/command"
	"github.com/percipia/eslgo/command/call"
)

type BridgeCommand struct {
	UId

	otherLeg interface{}
}

func (a *BridgeCommand) Validate() error {
	if err := a.UId.Validate(); err != nil {
		return err
	}

	if a.otherLeg == nil {
		return fmt.Errorf("other leg is required")
	}

	switch a.otherLeg.(type) {
	case *types.Leg:
		l := a.otherLeg.(*types.Leg)
		if !l.Valid() {
			return fmt.Errorf("other leg is not valid")
		}
	case string:
		if _, err := uuid.Parse(a.otherLeg.(string)); err != nil {
			return fmt.Errorf("invalid other leg uid: %v", err)
		}
	default:
		return fmt.Errorf("only supported other leg types are *types.Leg and uid")
	}

	return nil
}

func (a *BridgeCommand) Raw() (string, error) {
	if err := a.Validate(); err != nil {
		return "", err
	}

	var cmd command.Command
	if l, ok := a.otherLeg.(*types.Leg); ok {
		cmd = &call.Execute{UUID: a.uid, AppName: "bridge", AppArgs: l.DialString()}
	} else {
		cmd = &command.API{Command: "uuid_bridge", Arguments: fmt.Sprintf("%s %s", a.uid, a.otherLeg)}
	}

	return cmd.BuildMessage(), nil
}

func NewBridgeCommand(uid string, otherLeg *types.Leg) *BridgeCommand {
	return &BridgeCommand{UId: UId{uid: uid}, otherLeg: otherLeg}
}

var _ types.Command = (*BridgeCommand)(nil)
