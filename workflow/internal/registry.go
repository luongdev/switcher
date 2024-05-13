package internal

import "github.com/luongdev/switcher/workflow/types"

type RegistryImpl struct {
	activities map[string]types.Activity
	workflows  map[string]types.Workflow
}

func (r *RegistryImpl) RegisterWorkflow(name string, w types.Workflow) {

}

func (r *RegistryImpl) RegisterActivity(name string, a types.Activity) {
}

func (r *RegistryImpl) GetWorkflow(name string) (types.Workflow, bool) {
	//TODO implement me
	panic("implement me")
}

func (r *RegistryImpl) GetActivity(name string) (types.Activity, bool) {
	//TODO implement me
	panic("implement me")
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
