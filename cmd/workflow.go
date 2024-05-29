package main

import (
	"context"
	"encoding/json"
	"fmt"
	types2 "github.com/luongdev/switcher/freeswitch/types"
	"github.com/luongdev/switcher/internal/activities"
	pkg2 "github.com/luongdev/switcher/pkg"
	"github.com/luongdev/switcher/types"
	workflowtypes "github.com/luongdev/switcher/workflow/types"
	"go.uber.org/cadence/.gen/go/shared"
	libworkflow "go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"time"
)

//func main() {
//	cc := workflow.ClientConfig{
//		Host:       "103.141.141.60",
//		Port:       7833,
//		ClientName: "demo-client",
//	}
//
//	client, err := cc.Build()
//	if err != nil {
//		panic(err)
//	}
//
//	wc := workflow.WorkerConfig{
//		TaskList: "demo-task-list1",
//		Domains:  []string{"default"},
//	}
//
//	r := pkg.NewRegistry()
//	r.RegisterWorkflow("demo-workflow", &WorkflowImpl{registry: r})
//	r.RegisterActivity(pkg2.ActivityInitialize, internalactivity.NewInitializeActivity())
//
//	ws, err := wc.Build(client, r)
//	if err != nil {
//		panic(err)
//	}
//
//	if len(ws) > 0 {
//		go ws[0].Start()
//	}
//
//	go func() {
//		time.Sleep(10 * time.Second)
//
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		defer cancel()
//
//		uid := uuid.New().String()
//
//		executionStartToCloseTimeoutSeconds := int32(64)
//		taskStartToCloseTimeoutSeconds := int32(64)
//		n := "demo-workflow"
//
//		// start workflow
//		//wi := types.WorkflowInput{}
//		wo, err := client.StartWorkflowExecution(ctx, &shared.StartWorkflowExecutionRequest{
//			Domain:                              &wc.Domains[0],
//			WorkflowId:                          &uid,
//			RequestId:                           &uid,
//			WorkflowType:                        &shared.WorkflowType{Name: &n},
//			TaskList:                            &shared.TaskList{Name: &wc.TaskList},
//			Input:                               []byte("{}"),
//			ExecutionStartToCloseTimeoutSeconds: &executionStartToCloseTimeoutSeconds,
//			TaskStartToCloseTimeoutSeconds:      &taskStartToCloseTimeoutSeconds,
//		})
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("workflow started: %v\n", wo.GetRunId())
//	}()
//
//	select {}
//}

type WorkflowImpl struct {
	registry workflowtypes.Registry
	provider types2.ClientProvider
	client   workflowtypes.Client
}

var InboundSignal = "inbound"

func (w *WorkflowImpl) HandlerFunc() workflowtypes.WorkflowFunc {
	return func(ctx libworkflow.Context, input *workflowtypes.WorkflowInput) (o *workflowtypes.WorkflowOutput, err error) {

		logger := libworkflow.GetLogger(ctx)
		if a, ok := w.registry.GetActivity(pkg2.ActivityBridge); ok {
			sid := input.GetSessionId()
			if sid == "" {
				err = fmt.Errorf("sessionId is required")
				return
			}
			ai := activities.BridgeActivityInput{
				SessionId: sid,
				OtherLeg:  "sofia/develop/AGENT_10008@103.141.141.55:5080",
			}

			aCtx := libworkflow.WithActivityOptions(ctx, workflowtypes.ActivityTimeoutOptions(nil, time.Second*30))
			var o1 workflowtypes.ActivityOutput
			if err = libworkflow.ExecuteActivity(aCtx, a.HandlerFunc(), ai).Get(ctx, &o1); err != nil {
				return
			}

			fmt.Printf("activity output: %v\n", o1)
		}

		signalChan := libworkflow.GetSignalChannel(ctx, InboundSignal)
		defer signalChan.Close()

		var ws workflowtypes.WorkflowSignal
		var raw workflowtypes.Map
		sel := libworkflow.NewSelector(ctx).AddReceive(signalChan, func(ch libworkflow.Channel, ok bool) {
			if ok {
				ch.Receive(ctx, &raw)
				if err = raw.Convert(&ws); err != nil {
					logger.Error("Failed to convert signal", zap.Any("signal", raw), zap.Error(err))
				}
			}
		})

		go func(sid string) {
			<-time.NewTimer(time.Second * 30).C

			domain := "default"
			s := workflowtypes.WorkflowSignal{
				Action: pkg2.ActivityHangup,
				Input: map[string]interface{}{
					"sessionId":    sid,
					"hangupCause":  "normal_clearing",
					"hangupReason": "InboundSignalHangup",
				},
			}
			b, err := json.Marshal(s)
			if err != nil {
				logger.Error("Failed to marshal input", zap.Error(err))
				return
			}
			req := shared.SignalWorkflowExecutionRequest{
				Domain:            &domain,
				SignalName:        &InboundSignal,
				WorkflowExecution: &shared.WorkflowExecution{WorkflowId: &sid},
				Input:             b,
			}

			cc, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()
			err = w.client.SignalWorkflowExecution(cc, &req)
			if err != nil {
				return
			}
		}(input.GetSessionId())

		for {
			sel.Select(ctx)

			ws, err = ws.Default()
			if err != nil {
				logger.Error("Failed to process signal", zap.Any("signal", ws), zap.Error(err))
				continue
			}

			if act, ok := w.registry.GetActivity(ws.Action); ok {
				o := make(workflowtypes.Map)
				aCtx := libworkflow.WithActivityOptions(ctx, ws.Options(nil))
				err := libworkflow.ExecuteActivity(aCtx, act.HandlerFunc(), ws.Input).Get(aCtx, &o)
				if err != nil {
					logger.Error("Failed to execute activity", zap.Any("activity", ws.Action), zap.Error(err))
				} else {
					logger.Info("Activity executed", zap.Any("activity", ws.Action), zap.Any("output", o))
					ws = workflowtypes.WorkflowSignal{}
				}

				continue
			}
		}

		//if a, ok := w.registry.GetActivity(pkg2.ActivityInitialize); ok {
		//	ai := internalactivity.InitializeActivityInput{
		//		Initializer: "https://reqres.in/api/users?page=2",
		//		Protocol:    "http",
		//		Domain:      "voice.metechvn.com",
		//		ANI:         "0987654321",
		//		DNIS:        "0987654321",
		//		SessionId:   "dkalfja;klsdfja;sdjf",
		//	}
		//
		//	var o1 workflowtypes.ActivityOutput
		//	if err = libworkflow.ExecuteActivity(ctx, a.HandlerFunc(), ai).Get(ctx, &o1); err != nil {
		//		return
		//	}
		//
		//	if o1.Success && o1.Next != "" {
		//		if a, ok := w.registry.GetActivity(o1.Next); ok {
		//			if o1.Next == enums.ActivityHttp {
		//				hi := o1.Metadata[enums.FieldInput]
		//				if err = libworkflow.ExecuteActivity(ctx, a.HandlerFunc(), hi).Get(ctx, &o1); err != nil {
		//					return
		//				}
		//			}
		//		}
		//	}
		//
		//	fmt.Printf("activity output: %v\n", o)
		//
		//	return
		//}
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
