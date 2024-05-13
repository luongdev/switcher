package workflow

import (
	"fmt"
	"github.com/luongdev/switcher/workflow/internal"
	"github.com/luongdev/switcher/workflow/types"
	apiv1 "github.com/uber/cadence-idl/go/proto/api/v1"
	"go.uber.org/cadence/compatibility"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/grpc"
	"go.uber.org/yarpc/transport/tchannel"
)

type WorkerConfig struct {
	TaskList string   `yaml:"task_list"`
	Domains  []string `yaml:"domains"`
}

type ClientConfig struct {
	Host        string `yaml:"host"`
	Port        uint16 `yaml:"port"`
	ClientName  string `yaml:"client_name"`
	ServiceName string `yaml:"service_name"`
}

func (c *ClientConfig) Build() (types.Client, error) {
	if c.ServiceName == "" {
		c.ServiceName = "cadence-frontend"
	}
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == 0 {
		c.Port = 7833
	}

	if c.ClientName == "" {
		return nil, fmt.Errorf("client name is required")
	}

	hp := fmt.Sprintf("%v:%v", c.Host, c.Port)
	_, err := tchannel.NewChannelTransport(tchannel.ServiceName(c.ServiceName))
	if err != nil {
		return nil, err
	}

	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: c.ClientName,
		//Outbounds: yarpc.Outbounds{
		//	"cadence-frontend": {Unary: tChanTransport.NewSingleOutbound(hp)},
		//},
		Outbounds: yarpc.Outbounds{
			c.ServiceName: {Unary: grpc.NewTransport().NewSingleOutbound(hp)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		return nil, err
	}

	clientCfg := dispatcher.ClientConfig(c.ServiceName)
	itf := compatibility.NewThrift2ProtoAdapter(
		apiv1.NewDomainAPIYARPCClient(clientCfg),
		apiv1.NewWorkflowAPIYARPCClient(clientCfg),
		apiv1.NewWorkerAPIYARPCClient(clientCfg),
		apiv1.NewVisibilityAPIYARPCClient(clientCfg),
	)

	return internal.NewClient(itf, c.ClientName), nil
}

func (c *WorkerConfig) Build(client types.Client, registry types.Registry) ([]types.Worker, error) {
	if c.TaskList == "" {
		return nil, fmt.Errorf("task list is required")
	}

	logger, err := internal.NewLogger()
	if err != nil {
		return nil, err
	}

	workers := make([]types.Worker, 0, len(c.Domains))
	for _, domain := range c.Domains {
		if domain == "" {
			continue
		}

		worker := internal.NewWorker(domain, c.TaskList, client, registry, logger)
		workers = append(workers, worker)
	}

	if len(workers) == 0 {
		return nil, fmt.Errorf("no domains provided")
	}

	return workers, nil
}
