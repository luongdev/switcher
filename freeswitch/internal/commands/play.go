package commands

import (
	"fmt"
	"github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/freeswitch/utils"
	"github.com/percipia/eslgo/command/call"
)

type PlayCommand struct {
	UId

	path string
}

func (p *PlayCommand) Validate() error {
	if err := p.UId.Validate(); err != nil {
		return err
	}

	if !utils.IsPathValid(p.path, ".wav") {
		return fmt.Errorf("path must be a .wav file. Got %v", p.path)
	}

	return nil
}

func (p *PlayCommand) Raw() (string, error) {
	if err := p.Validate(); err != nil {
		return "", err
	}

	return (&call.Execute{UUID: p.uid, AppName: "playback", AppArgs: p.path}).BuildMessage(), nil
}

func NewPlayCommand(uid, path string) *PlayCommand {
	return &PlayCommand{UId: UId{uid: uid}, path: path}
}

var _ types.Command = (*PlayCommand)(nil)
