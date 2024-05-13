package main

import (
	"github.com/luongdev/switcher/workflow"
	"github.com/luongdev/switcher/workflow/pkg"
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
		TaskList: "demo-task-list",
		Domains:  []string{"default"},
	}

	r := pkg.NewRegistry()

	ws, err := wc.Build(client, r)
	if err != nil {
		panic(err)
	}

	if len(ws) > 0 {
		err = ws[0].Start()
		if err != nil {
			panic(err)
		}
	}
}
