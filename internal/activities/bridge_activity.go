package activities

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/luongdev/switcher/freeswitch/pkg"
	freeswitchtypes "github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal"
	"github.com/luongdev/switcher/types"
	"github.com/luongdev/switcher/workflow/enums"
	workflowtypes "github.com/luongdev/switcher/workflow/types"
)

type BridgeActivityInput struct {
	SessionId string `input:"sessionId" json:"sessionId"`
	OtherLeg  string `input:"otherLeg" json:"otherLeg"`

	otherLeg interface{}
}

func (n *BridgeActivityInput) GetSessionId() string {
	return n.SessionId
}

var _ types.SessionInput = (*BridgeActivityInput)(nil)

func (n *BridgeActivityInput) ValidateAndDefault() error {
	if n.GetSessionId() == "" {
		return fmt.Errorf("sessionId is required")
	}

	if n.OtherLeg == "" {
		return fmt.Errorf("otherLeg is required")
	}

	if _, err := uuid.Parse(n.OtherLeg); err == nil {
		n.otherLeg = n.OtherLeg

		return nil
	}

	n.otherLeg = &freeswitchtypes.Leg{Endpoint: n.OtherLeg}

	return nil
}

type BridgeActivityOutput struct {
}

type BridgeActivity struct {
	internal.SessionActivity

	input BridgeActivityInput
}

func (b *BridgeActivity) HandlerFunc() workflowtypes.ActivityFunc {
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

func (b *BridgeActivity) execute(ctx context.Context) (o *workflowtypes.ActivityOutput, err error) {
	o = &workflowtypes.ActivityOutput{Metadata: make(map[enums.Field]interface{})}
	_ = &BridgeActivityOutput{}

	client, ok := b.GetClient(&b.input)
	if !ok {
		err = fmt.Errorf("failed to get client for session %s", b.input.GetSessionId())
		return
	}
	res, err := client.Exec(ctx, pkg.BridgeCommand(b.input.GetSessionId(), b.input.otherLeg))
	if err != nil {
		return nil, err
	}

	o.Metadata[enums.FieldOutput] = res

	return
}

func NewBridgeActivity(provider freeswitchtypes.ClientProvider) *BridgeActivity {
	return &BridgeActivity{SessionActivity: internal.NewFreeswitchActivity(provider)}
}

var _ workflowtypes.Activity = (*BridgeActivity)(nil)
