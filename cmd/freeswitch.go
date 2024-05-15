package main

import (
	"context"
	"fmt"
	"github.com/luongdev/switcher/freeswitch"
	"github.com/luongdev/switcher/freeswitch/pkg"
	"github.com/luongdev/switcher/freeswitch/types"
	internalactivity "github.com/luongdev/switcher/internal/activities"
	pkg2 "github.com/luongdev/switcher/pkg"
	"github.com/luongdev/switcher/workflow"
	pkg3 "github.com/luongdev/switcher/workflow/pkg"
	types2 "github.com/luongdev/switcher/workflow/types"
	"go.uber.org/cadence/.gen/go/shared"
	"log"
	"time"
)

func main() {
	c := freeswitch.InboundConfig{
		Host:           "10.8.0.1",
		Port:           65021,
		Password:       "Simplefs!!",
		ConnectTimeout: time.Millisecond * 2,
	}

	store := pkg.NewClientStore(nil)

	fsClient, err := c.Build()
	store.Set("", fsClient)

	if err != nil {
		panic(err)
	}

	co := freeswitch.OutboundConfig{
		Host: "10.8.0.2",
		Port: 65022,
	}

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

	r := pkg3.NewRegistry()
	provider := pkg.NewClientProvider(store)

	r.RegisterWorkflow("demo-workflow", &WorkflowImpl{registry: r, provider: provider})
	r.RegisterActivity(pkg2.ActivitySessionInit, internalactivity.NewNewSessionActivity())
	r.RegisterActivity(pkg2.ActivityBridge, internalactivity.NewBridgeActivity(provider))

	server := co.Build()
	server.SetStore(store)
	server.OnSessionStarted(func(ctx context.Context, session types.Session) {
		log.Printf("Session started: %s", session.GetId())

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		executionStartToCloseTimeoutSeconds := int32(64)
		taskStartToCloseTimeoutSeconds := int32(64)
		n := "demo-workflow"

		sid := session.GetId()
		m := types2.Map{"sessionId": sid}
		ib, err := m.Bytes()
		if err != nil {
			panic(err)
		}

		// start workflow
		//wi := types.WorkflowInput{}
		wo, err := client.StartWorkflowExecution(ctx, &shared.StartWorkflowExecutionRequest{
			Domain:                              &wc.Domains[0],
			WorkflowId:                          &sid,
			RequestId:                           &sid,
			WorkflowType:                        &shared.WorkflowType{Name: &n},
			TaskList:                            &shared.TaskList{Name: &wc.TaskList},
			Input:                               ib,
			ExecutionStartToCloseTimeoutSeconds: &executionStartToCloseTimeoutSeconds,
			TaskStartToCloseTimeoutSeconds:      &taskStartToCloseTimeoutSeconds,
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("workflow started: %v\n", wo.GetRunId())

		//bridgeCmd := pkg.NewBridgeCommand(session.GetId(), &types.Leg{
		//	Endpoint: "sofia/external/AGENT_10008@103.141.141.55:5080",
		//})
		bridgeCmd := pkg.LuaCommand(session.GetId(), "default_ivr.lua")
		out, err := session.Exec(ctx, bridgeCmd)
		if err != nil {
			log.Printf("Failed to execute bridge command: %s", err)
			return
		}

		log.Printf("Bridge command response: %s", out)

		//_, err = session.Exec(ctx, pkg.SetCommand(session.GetId(), map[string]interface{}{
		//	"effective_caller_id_name":     "Test",
		//	"effective_caller_id_number":   "1234567890",
		//	"origination_caller_id_name":   "Test",
		//	"origination_caller_id_number": "1234567890",
		//	"origination_uuid":             "1234567890",
		//}))
		//
		//if err != nil {
		//	log.Printf("Failed to set session variables: %s", err)
		//	return
		//}
		//
		//if err := session.Answer(ctx); err != nil {
		//	log.Printf("Failed to answer session: %s", err)
		//	return
		//}
		//
		//log.Printf("Answered session: %s", session.GetId())
		//
		//origCmd := pkg.NewOriginateCommand(false, &types.Leg{
		//	Endpoint: "sofia/external/TO_IVR@103.141.141.55:5080",
		//	Uid:      uuid.New().String(),
		//}, &types.Leg{Endpoint: "&sleep(30000)"}, nil)
		//
		//if res, err := session.Exec(ctx, origCmd); err != nil {
		//	log.Printf("Failed to execute originate command: %s", err)
		//	return
		//} else {
		//	log.Printf("Originate command response: %s", res)
		//}
		//
		//if err := session.Hangup(ctx, "CALL_REJECTED"); err != nil {
		//	log.Printf("Failed to hangup session: %s", err)
		//	return
		//}
		//
		//log.Printf("Hung up session: %s", session.GetId())

	})

	go func() {
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()

	ws, err := wc.Build(client, r)
	if err != nil {
		panic(err)
	}

	if len(ws) > 0 {
		w := ws[0]
		go w.Start()
	}

	select {}
}
