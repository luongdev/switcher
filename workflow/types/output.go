package types

import "github.com/luongdev/switcher/workflow/enums"

type WorkflowOutput struct {
	Success  bool     `json:"success"`
	Metadata Metadata `json:"metadata"`
}

type ActivityOutput struct {
	Success  bool           `json:"success"`
	Next     enums.Activity `json:"next"`
	Metadata Metadata       `json:"metadata"`
}
