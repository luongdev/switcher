package types

type Registry interface {
	RegisterWorkflow(name string, w Workflow)
	RegisterActivity(name string, a Activity)

	GetWorkflow(name string) (Workflow, bool)
	GetActivity(name string) (Activity, bool)
}
