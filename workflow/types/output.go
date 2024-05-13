package types

import "github.com/luongdev/switcher/workflow/enums"

type WorkflowOutput struct {
	Success  bool `json:"success"`
	Metadata map[enums.Field]interface{}
}

type ActivityOutput struct {
	Success  bool                        `json:"success"`
	Next     enums.Activity              `json:"next"`
	Metadata map[enums.Field]interface{} `json:"metadata"`
}
