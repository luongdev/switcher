package types

import "github.com/luongdev/switcher/workflow/enums"

type WorkflowOutput struct {
	Success  bool `json:"success"`
	Metadata map[enums.Field]interface{}
}

type ActivityOutput struct {
	Success  bool `json:"success"`
	Metadata map[enums.Field]interface{}
}
