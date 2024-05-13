package internal

import (
	"github.com/luongdev/switcher/workflow/enums"
	"github.com/luongdev/switcher/workflow/pkg/activities"
	"github.com/luongdev/switcher/workflow/types"
)

type RegistryImpl struct {
	activities map[enums.Activity]types.Activity
	workflows  map[string]types.Workflow
}

func (r *RegistryImpl) RegisterWorkflow(name string, w types.Workflow) {
	if name == "" {
		return
	}

	if _, ok := r.workflows[name]; ok {
		return
	}

	r.workflows[name] = w
}

func (r *RegistryImpl) RegisterActivity(name enums.Activity, a types.Activity) {
	if name == "" {
		return
	}

	if _, ok := r.activities[name]; ok {
		return
	}

	r.activities[name] = a
}

func (r *RegistryImpl) GetWorkflow(name string) (types.Workflow, bool) {
	if w, ok := r.workflows[name]; ok {
		return w, true
	}

	return nil, false
}

func (r *RegistryImpl) GetActivity(name enums.Activity) (types.Activity, bool) {
	if a, ok := r.activities[name]; ok {
		return a, true
	}

	return nil, false
}

func (r *RegistryImpl) Workflows() map[string]types.Workflow {
	return r.workflows
}

func (r *RegistryImpl) Activities() map[enums.Activity]types.Activity {
	return r.activities
}

func NewRegistry() types.Registry {
	r := &RegistryImpl{
		workflows:  make(map[string]types.Workflow),
		activities: make(map[enums.Activity]types.Activity),
	}

	r.RegisterActivity(enums.ActivityHttp, activities.HttpActivity())

	return r
}

var _ types.Registry = (*RegistryImpl)(nil)
