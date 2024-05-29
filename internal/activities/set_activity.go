package activities

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/freeswitch/pkg"
	freeswitchtypes "github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal"
	"github.com/luongdev/switcher/types"
	"github.com/luongdev/switcher/workflow/enums"
	workflowtypes "github.com/luongdev/switcher/workflow/types"
)

type SetActivityInput struct {
	SessionId string                 `input:"sessionId" json:"sessionId"`
	Variables map[string]interface{} `input:"variables" json:"variables"`
}

func (n *SetActivityInput) GetSessionId() string {
	return n.SessionId
}

var _ types.SessionInput = (*SetActivityInput)(nil)

func (n *SetActivityInput) ValidateAndDefault() error {
	if n.GetSessionId() == "" {
		return fmt.Errorf("sessionId is required")
	}

	if len(n.Variables) == 0 {
		return fmt.Errorf("variables is required")
	}

	return nil
}

type SetActivityOutput struct {
}

type SetActivity struct {
	internal.SessionActivity

	input SetActivityInput
}

func (b *SetActivity) HandlerFunc() workflowtypes.ActivityFunc {
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

func (b *SetActivity) execute(ctx context.Context) (o *workflowtypes.ActivityOutput, err error) {
	o = &workflowtypes.ActivityOutput{Metadata: make(map[enums.Field]interface{})}
	_ = &SetActivityOutput{}

	client, ok := b.GetClient(&b.input)
	if !ok {
		err = fmt.Errorf("failed to get client for session %s", b.input.GetSessionId())
		return
	}
	res, err := client.Exec(ctx, pkg.SetCommand(b.input.GetSessionId(), b.input.Variables))
	if err != nil {
		return nil, err
	}

	o.Metadata[enums.FieldOutput] = res

	return
}

func NewSetActivity(provider freeswitchtypes.ClientProvider) *SetActivity {
	return &SetActivity{SessionActivity: internal.NewFreeswitchActivity(provider)}
}

var _ workflowtypes.Activity = (*SetActivity)(nil)
