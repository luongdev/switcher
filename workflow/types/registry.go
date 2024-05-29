package types

import "github.com/luongdev/switcher/workflow/enums"

type Registry interface {
	RegisterWorkflow(name string, w Workflow)
	RegisterActivity(name enums.Activity, a Activity)

	GetWorkflow(name string) (Workflow, bool)
	GetActivity(name enums.Activity) (Activity, bool)
}
