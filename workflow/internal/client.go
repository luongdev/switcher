package internal

import (
	"github.com/luongdev/switcher/workflow/types"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
)

type ClientImpl struct {
	workflowserviceclient.Interface

	clientName string
}

func (c *ClientImpl) GetName() string {
	return c.clientName
}

func NewClient(client workflowserviceclient.Interface, clientName string) types.Client {
	return &ClientImpl{
		Interface:  client,
		clientName: clientName,
	}
}

var _ types.Client = (*ClientImpl)(nil)
