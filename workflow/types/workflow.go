package types

import "go.uber.org/cadence/workflow"

type WorkflowFunc func(workflow.Context, *WorkflowInput) (*WorkflowOutput, error)

type Workflow interface {
	HandlerFunc() WorkflowFunc
}
