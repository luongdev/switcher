package types

import (
	"fmt"
	"github.com/luongdev/switcher/workflow/enums"
	"go.uber.org/cadence/workflow"
	"time"
)

type WorkflowInput Map

func (i WorkflowInput) GetSessionId() string {
	if id, ok := i["sessionId"]; ok {
		return fmt.Sprintf("%v", id)
	}

	return ""
}

func (i WorkflowInput) Convert(o interface{}) error {
	return Map(i).Convert(o)
}

type ActivityInput Map

func (i ActivityInput) Timeout() time.Duration {
	timeout := time.Minute
	if t, ok := i["timeout"]; ok {
		if d, ok := t.(time.Duration); ok {
			timeout = d
		}
	}

	return timeout
}

func (i ActivityInput) Callback() string {
	callback := ""
	if c, ok := i["callback"]; ok {
		if s, ok := c.(string); ok {
			callback = s
		}
	}

	return callback
}

func (i ActivityInput) Options(parent *workflow.ActivityOptions) workflow.ActivityOptions {
	return ActivityTimeoutOptions(parent, i.Timeout())
}

func (i ActivityInput) Convert(o interface{}) error {
	return Map(i).Convert(o)
}

type WorkflowSignal struct {
	Action  enums.Activity `json:"action"`
	Input   ActivityInput  `json:"input"`
	Timeout time.Duration  `json:"timeout"`
}

func (s WorkflowSignal) Default() (WorkflowSignal, error) {
	if s.Action == "" {
		return s, fmt.Errorf("action is required")
	}

	if s.Input == nil {
		s.Input = make(ActivityInput)
	}

	if s.Timeout <= 0 {
		s.Timeout = time.Minute
	}

	return s, nil
}

func (s WorkflowSignal) Options(parent *workflow.ActivityOptions) workflow.ActivityOptions {
	return ActivityTimeoutOptions(parent, s.Timeout)
}

func ActivityTimeoutOptions(parent *workflow.ActivityOptions, timeout time.Duration) workflow.ActivityOptions {
	opts := &workflow.ActivityOptions{}
	if parent != nil {
		opts = parent
	} else if timeout <= 0 {
		timeout = time.Minute
	}

	opts.StartToCloseTimeout = timeout
	if opts.HeartbeatTimeout <= 0 {
		opts.HeartbeatTimeout = timeout / 5
	}
	if opts.ScheduleToStartTimeout <= 0 {
		opts.ScheduleToStartTimeout = opts.HeartbeatTimeout
	}

	return *opts
}
