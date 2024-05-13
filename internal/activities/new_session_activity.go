package activities

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/types"
	"github.com/luongdev/switcher/workflow/enums"
	"github.com/luongdev/switcher/workflow/pkg/activities"
	workflowtypes "github.com/luongdev/switcher/workflow/types"
	"net/http"
	"time"
)

type NewSessionActivityInput struct {
	ANI         string        `input:"ani" json:"ani"`
	DNIS        string        `input:"dnis" json:"dnis"`
	Domain      string        `input:"domain" json:"domain"`
	Initializer string        `input:"initializer" json:"initializer"`
	Protocol    string        `input:"protocol" json:"protocol"`
	Timeout     time.Duration `input:"timeout" json:"timeout"`
	SessionId   string        `input:"sessionId" json:"sessionId"`
}

func (n *NewSessionActivityInput) GetSessionId() string {
	return n.SessionId
}

var _ types.SessionInput = (*NewSessionActivityInput)(nil)

func (n *NewSessionActivityInput) ValidateAndDefault() error {
	if n.GetSessionId() == "" {
		return fmt.Errorf("sessionId is required")
	}
	if n.Timeout <= time.Millisecond*100 {
		n.Timeout = time.Millisecond * 1500
	}
	if n.Domain == "" {
		return fmt.Errorf("domain is required")
	}
	if n.Initializer == "" {
		return fmt.Errorf("initializer is required")
	}
	if n.ANI == "" {
		return fmt.Errorf("ani is required")
	}
	if n.DNIS == "" {
		return fmt.Errorf("dnis is required")
	}
	if n.Protocol == "" {
		n.Protocol = "http"
	}

	return nil
}

type NewSessionActivityOutput struct {
}

type NewSessionActivity struct {
	input NewSessionActivityInput
}

func (n *NewSessionActivity) HandlerFunc() workflowtypes.ActivityFunc {
	return func(ctx context.Context, i *workflowtypes.ActivityInput) (o *workflowtypes.ActivityOutput, err error) {
		if err = i.Convert(&n.input); err != nil {
			return
		}

		if err = n.input.ValidateAndDefault(); err != nil {
			return
		}

		return n.execute()
	}
}

func (n *NewSessionActivity) execute() (o *workflowtypes.ActivityOutput, err error) {
	o = &workflowtypes.ActivityOutput{Metadata: make(map[enums.Field]interface{})}
	ao := &NewSessionActivityOutput{}
	switch n.input.Protocol {
	default:
		o.Success = true
		o.Next = enums.ActivityHttp
		o.Metadata[enums.FieldSessionId] = n.input.GetSessionId()
		o.Metadata[enums.FieldOutput] = ao

		fi := &activities.HttpActivityInput{
			Url:     n.input.Initializer,
			Timeout: n.input.Timeout,
			Method:  http.MethodPost,
			Headers: workflowtypes.Map{"Domain-Name": n.input.Domain},
			Body:    workflowtypes.Map{"ani": n.input.ANI, "dnis": n.input.DNIS, "sessionId": n.input.GetSessionId()},
		}

		o.Metadata[enums.FieldInput] = fi
	}
	return
}

func NewNewSessionActivity() *NewSessionActivity {
	return &NewSessionActivity{}
}

var _ workflowtypes.Activity = (*NewSessionActivity)(nil)
