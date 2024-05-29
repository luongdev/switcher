package activities

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/freeswitch/pkg"
	freeswitchtypes "github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/freeswitch/utils"
	"github.com/luongdev/switcher/internal"
	"github.com/luongdev/switcher/types"
	"github.com/luongdev/switcher/workflow/enums"
	workflowtypes "github.com/luongdev/switcher/workflow/types"
)

type LuaActivityInput struct {
	SessionId string `input:"sessionId" json:"sessionId"`
	Path      string `input:"path" json:"path"`
}

func (n *LuaActivityInput) GetSessionId() string {
	return n.SessionId
}

var _ types.SessionInput = (*LuaActivityInput)(nil)

func (n *LuaActivityInput) ValidateAndDefault() error {
	if !utils.IsPathValid(n.Path, ".lua") {
		return fmt.Errorf("path must be a .lua file. Got %v", n.Path)
	}

	return nil
}

type LuaActivityOutput struct {
}

type LuaActivity struct {
	internal.SessionActivity

	input LuaActivityInput
}

func (b *LuaActivity) HandlerFunc() workflowtypes.ActivityFunc {
	return func(ctx context.Context, i *workflowtypes.ActivityInput) (o *workflowtypes.ActivityOutput, err error) {
		if err = i.Convert(&b.input); err != nil {
			return
		}

		if err = b.input.ValidateAndDefault(); err != nil {
			return
		}

		return b.execute(ctx)
	}
}

func (b *LuaActivity) execute(ctx context.Context) (o *workflowtypes.ActivityOutput, err error) {
	o = &workflowtypes.ActivityOutput{Metadata: make(map[enums.Field]interface{})}
	_ = &LuaActivityOutput{}

	client, ok := b.GetClient(&b.input)
	if !ok {
		err = fmt.Errorf("failed to get client for session %s", b.input.GetSessionId())
		return
	}
	res, err := client.Exec(ctx, pkg.LuaCommand(b.input.GetSessionId(), b.input.Path))
	if err != nil {
		return nil, err
	}

	o.Metadata[enums.FieldOutput] = res

	return
}

func NewLuaActivity(provider freeswitchtypes.ClientProvider) *LuaActivity {
	return &LuaActivity{SessionActivity: internal.NewFreeswitchActivity(provider)}
}

var _ workflowtypes.Activity = (*LuaActivity)(nil)
