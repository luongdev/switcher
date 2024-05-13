package internal

import "github.com/luongdev/switcher/workflow/types"

type RegistryImpl struct {
	activities map[string]types.Activity
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

func (r *RegistryImpl) RegisterActivity(name string, a types.Activity) {
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

func (r *RegistryImpl) GetActivity(name string) (types.Activity, bool) {
	if a, ok := r.activities[name]; ok {
		return a, true
	}

	return nil, false
}

func (r *RegistryImpl) Workflows() map[string]types.Workflow {
	return r.workflows
}

func (r *RegistryImpl) Activities() map[string]types.Activity {
	return r.activities
}

func NewRegistry() types.Registry {
	return &RegistryImpl{
		activities: make(map[string]types.Activity),
		workflows:  make(map[string]types.Workflow),
	}
}

var _ types.Registry = (*RegistryImpl)(nil)
