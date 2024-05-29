package types

import "go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"

type Client interface {
	workflowserviceclient.Interface

	GetName() string
}
