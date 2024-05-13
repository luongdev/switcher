package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	internalactivity "github.com/luongdev/switcher/internal/activities"
	pkg2 "github.com/luongdev/switcher/pkg"
	"github.com/luongdev/switcher/types"
	"github.com/luongdev/switcher/workflow"
	"github.com/luongdev/switcher/workflow/enums"
	"github.com/luongdev/switcher/workflow/pkg"
	workflowtypes "github.com/luongdev/switcher/workflow/types"
	"go.uber.org/cadence/.gen/go/shared"
	libworkflow "go.uber.org/cadence/workflow"
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
	r.RegisterActivity(pkg2.ActivitySessionInit, internalactivity.NewNewSessionActivity())

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
	registry workflowtypes.Registry
}

func (w *WorkflowImpl) HandlerFunc() workflowtypes.WorkflowFunc {
	return func(ctx libworkflow.Context, input *workflowtypes.WorkflowInput) (o *workflowtypes.WorkflowOutput, err error) {

		ctx = libworkflow.WithActivityOptions(ctx,
			libworkflow.ActivityOptions{ScheduleToStartTimeout: time.Second, StartToCloseTimeout: time.Second * 60})

		if a, ok := w.registry.GetActivity(pkg2.ActivitySessionInit); ok {
			ai := internalactivity.NewSessionActivityInput{
				Initializer: "https://reqres.in/api/users?page=2",
				Protocol:    "http",
				Domain:      "voice.metechvn.com",
				ANI:         "0987654321",
				DNIS:        "0987654321",
				SessionId:   "dkalfja;klsdfja;sdjf",
			}

			var o1 workflowtypes.ActivityOutput
			if err = libworkflow.ExecuteActivity(ctx, a.HandlerFunc(), ai).Get(ctx, &o1); err != nil {
				return
			}

			if o1.Success && o1.Next != "" {
				if a, ok := w.registry.GetActivity(o1.Next); ok {
					if o1.Next == enums.ActivityHttp {
						hi := o1.Metadata[enums.FieldInput]
						if err = libworkflow.ExecuteActivity(ctx, a.HandlerFunc(), hi).Get(ctx, &o1); err != nil {
							return
						}
					}
				}
			}

			fmt.Printf("activity output: %v\n", o)

			return
		}

		return
	}
}

var _ workflowtypes.Workflow = (*WorkflowImpl)(nil)

type SS struct {
	SessionId string `json:"sessionId"`
}

func (s *SS) GetSessionId() string {
	return s.SessionId
}

var _ types.SessionInput = (*SS)(nil)
