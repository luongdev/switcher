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

type HangupActivityInput struct {
	SessionId    string `input:"sessionId" json:"sessionId"`
	HangupCause  string `input:"hangupCause" json:"hangupCause"`
	HangupReason string `input:"hangupReason" json:"hangupReason"`
}

func (n *HangupActivityInput) GetSessionId() string {
	return n.SessionId
}

var _ types.SessionInput = (*HangupActivityInput)(nil)

func (n *HangupActivityInput) ValidateAndDefault() error {
	if n.GetSessionId() == "" {
		return fmt.Errorf("sessionId is required")
	}

	return nil
}

type HangupActivityOutput struct {
}

type HangupActivity struct {
	internal.SessionActivity

	input HangupActivityInput
}

func (h *HangupActivity) HandlerFunc() workflowtypes.ActivityFunc {
	return func(ctx context.Context, i *workflowtypes.ActivityInput) (o *workflowtypes.ActivityOutput, err error) {
		if err = i.Convert(&h.input); err != nil {
			return
		}

		if err = h.input.ValidateAndDefault(); err != nil {
			return
		}

		return h.execute(ctx)
	}
}

func (h *HangupActivity) execute(ctx context.Context) (o *workflowtypes.ActivityOutput, err error) {
	o = &workflowtypes.ActivityOutput{Metadata: make(map[enums.Field]interface{})}
	_ = &HangupActivityOutput{}

	client, ok := h.GetClient(&h.input)
	if !ok {
		err = fmt.Errorf("failed to get client for session %s", h.input.GetSessionId())
		return
	}
	_, _ = client.Exec(ctx, pkg.SetCommand(
		h.input.GetSessionId(),
		map[string]interface{}{"hangup_reason": h.input.HangupReason},
	))
	res, err := client.Exec(ctx, pkg.HangupCommand(h.input.GetSessionId(), h.input.HangupCause))
	if err != nil {
		return nil, err
	}

	o.Metadata[enums.FieldOutput] = res

	return
}

func NewHangupActivity(provider freeswitchtypes.ClientProvider) *HangupActivity {
	return &HangupActivity{SessionActivity: internal.NewFreeswitchActivity(provider)}
}

var _ workflowtypes.Activity = (*HangupActivity)(nil)
