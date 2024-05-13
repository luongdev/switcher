package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/luongdev/switcher/workflow"
	"github.com/luongdev/switcher/workflow/pkg"
	"github.com/luongdev/switcher/workflow/pkg/activities"
	"github.com/luongdev/switcher/workflow/types"
	"go.uber.org/cadence/.gen/go/shared"
	libworkflow "go.uber.org/cadence/workflow"
	"net/http"
	"time"
)

func main() {
	cc := workflow.ClientConfig{
		Host:       "103.141.141.60",
		Port:       7833,
		ClientName: "demo-client",
	}

	client, err := cc.Build()
	if err != nil {
		panic(err)
	}

	wc := workflow.WorkerConfig{
		TaskList: "demo-task-list1",
		Domains:  []string{"default"},
	}

	r := pkg.NewRegistry()
	r.RegisterWorkflow("demo-workflow", &WorkflowImpl{registry: r})
	r.RegisterActivity("http", activities.HttpActivity())

	ws, err := wc.Build(client, r)
	if err != nil {
		panic(err)
	}

	if len(ws) > 0 {
		go ws[0].Start()
	}

	go func() {
		time.Sleep(10 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		uid := uuid.New().String()

		executionStartToCloseTimeoutSeconds := int32(64)
		taskStartToCloseTimeoutSeconds := int32(64)
		n := "demo-workflow"

		// start workflow
		//wi := types.WorkflowInput{}
		wo, err := client.StartWorkflowExecution(ctx, &shared.StartWorkflowExecutionRequest{
			Domain:                              &wc.Domains[0],
			WorkflowId:                          &uid,
			RequestId:                           &uid,
			WorkflowType:                        &shared.WorkflowType{Name: &n},
			TaskList:                            &shared.TaskList{Name: &wc.TaskList},
			Input:                               []byte("{}"),
			ExecutionStartToCloseTimeoutSeconds: &executionStartToCloseTimeoutSeconds,
			TaskStartToCloseTimeoutSeconds:      &taskStartToCloseTimeoutSeconds,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("workflow started: %v\n", wo.GetRunId())
	}()

	select {}
}

type WorkflowImpl struct {
	registry types.Registry
}

func (w *WorkflowImpl) HandlerFunc() types.WorkflowFunc {
	return func(ctx libworkflow.Context, input *types.WorkflowInput) (o *types.WorkflowOutput, err error) {

		ctx = libworkflow.WithActivityOptions(ctx,
			libworkflow.ActivityOptions{ScheduleToStartTimeout: time.Second, StartToCloseTimeout: time.Second * 60})

		if a, ok := w.registry.GetActivity("http"); ok {
			ai := activities.HttpActivityInput{
				Url:    "https://reqres.in/api/users?page=2",
				Method: http.MethodGet,
			}

			var o1 interface{}
			if err = libworkflow.ExecuteActivity(ctx, a.HandlerFunc(), ai).Get(ctx, &o1); err != nil {
				return
			}

			fmt.Printf("activity output: %v\n", o)

			return
		}

		return
	}
}

var _ types.Workflow = (*WorkflowImpl)(nil)
