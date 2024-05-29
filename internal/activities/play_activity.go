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

type PlayActivityInput struct {
	SessionId string `input:"sessionId" json:"sessionId"`
	Path      string `input:"path" json:"path"`
}

func (n *PlayActivityInput) GetSessionId() string {
	return n.SessionId
}

var _ types.SessionInput = (*PlayActivityInput)(nil)

func (n *PlayActivityInput) ValidateAndDefault() error {
	if n.GetSessionId() == "" {
		return fmt.Errorf("sessionId is required")
	}

	if !utils.IsPathValid(n.Path, ".wav") {
		return fmt.Errorf("path must be a .wav file. Got %v", n.Path)
	}

	return nil
}

type PlayActivityOutput struct {
}

type PlayActivity struct {
	internal.SessionActivity

	input PlayActivityInput
}

func (b *PlayActivity) HandlerFunc() workflowtypes.ActivityFunc {
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

func (b *PlayActivity) execute(ctx context.Context) (o *workflowtypes.ActivityOutput, err error) {
	o = &workflowtypes.ActivityOutput{Metadata: make(map[enums.Field]interface{})}
	_ = &PlayActivityOutput{}

	client, ok := b.GetClient(&b.input)
	if !ok {
		err = fmt.Errorf("failed to get client for session %s", b.input.GetSessionId())
		return
	}
	res, err := client.Exec(ctx, pkg.PlayCommand(b.input.GetSessionId(), b.input.Path))
	if err != nil {
		return nil, err
	}

	o.Metadata[enums.FieldOutput] = res

	return
}

func NewPlayActivity(provider freeswitchtypes.ClientProvider) *PlayActivity {
	return &PlayActivity{SessionActivity: internal.NewFreeswitchActivity(provider)}
}

var _ workflowtypes.Activity = (*PlayActivity)(nil)
