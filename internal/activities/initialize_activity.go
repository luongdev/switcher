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

type InitializeActivityInput struct {
	ANI         string        `input:"ani" json:"ani,abc,def"`
	DNIS        string        `input:"dnis" json:"dnis"`
	Domain      string        `input:"domain" json:"domain"`
	Initializer string        `input:"initializer" json:"initializer"`
	Protocol    string        `input:"protocol" json:"protocol"`
	Timeout     time.Duration `input:"timeout" json:"timeout"`
	SessionId   string        `input:"sessionId" json:"sessionId"`
}

func (n *InitializeActivityInput) GetSessionId() string {
	return n.SessionId
}

var _ types.SessionInput = (*InitializeActivityInput)(nil)

func (n *InitializeActivityInput) ValidateAndDefault() error {
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

type InitializeActivityOutput struct {
}

type InitializeActivity struct {
	input InitializeActivityInput
}

func (n *InitializeActivity) HandlerFunc() workflowtypes.ActivityFunc {
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

func (n *InitializeActivity) execute() (o *workflowtypes.ActivityOutput, err error) {
	o = &workflowtypes.ActivityOutput{Metadata: make(map[enums.Field]interface{})}
	ao := &InitializeActivityOutput{}
	switch n.input.Protocol {
	default:
		o.Success = true
		o.Next = enums.ActivityHttp
		o.Metadata[enums.FieldSessionId] = n.input.GetSessionId()
		o.Metadata[enums.FieldOutput] = ao

		body := workflowtypes.Map{}
		err = body.Set(&n.input)
		if err != nil {
			return
		}
		fi := &activities.HttpActivityInput{
			Url:     n.input.Initializer,
			Timeout: n.input.Timeout,
			Method:  http.MethodPost,
			Headers: workflowtypes.Map{"Domain-Name": n.input.Domain},
			Body:    body,
		}

		o.Metadata[enums.FieldInput] = fi
	}
	return
}

func NewInitializeActivity() *InitializeActivity {
	return &InitializeActivity{}
}

var _ workflowtypes.Activity = (*InitializeActivity)(nil)
