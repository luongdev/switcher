package pkg

import "github.com/luongdev/switcher/workflow/types"

type SessionInput struct {
	types.WorkflowInput `json:"-"`

	SessionId string `input:"sessionId" json:"sessionId"`
}
