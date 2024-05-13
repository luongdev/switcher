package types

type WorkflowFunc func(WorkflowInput) (WorkflowOutput, error)

type Workflow interface {
	HandlerFunc() WorkflowFunc
}
