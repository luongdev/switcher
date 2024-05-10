package main

import (
	"context"
	"github.com/luongdev/switcher/freeswitch"
	"github.com/luongdev/switcher/freeswitch/interfaces"
	"github.com/luongdev/switcher/freeswitch/pkg"
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

	client, err := c.Build()
	store.Set("", client)

	if err != nil {
		panic(err)
	}

	co := freeswitch.OutboundConfig{
		Host: "10.8.0.2",
		Port: 65022,
	}

	server := co.Build()
	server.SetStore(store)
	server.OnSessionStarted(func(ctx context.Context, session interfaces.Session) {
		log.Printf("Session started: %s", session.GetId())

		_, _ = session.Exec(ctx, pkg.SetCommand(session.GetId(), map[string]interface{}{
			"effective_caller_id_name":     "Test",
			"effective_caller_id_number":   "1234567890",
			"origination_caller_id_name":   "Test",
			"origination_caller_id_number": "1234567890",
			"origination_uuid":             "1234567890",
		}))

		if err := session.Answer(ctx); err != nil {
			log.Printf("Failed to answer session: %s", err)
			return
		}

		log.Printf("Answered session: %s", session.GetId())

		if err := session.Hangup(ctx, "CALL_REJECTED"); err != nil {
			log.Printf("Failed to hangup session: %s", err)
			return
		}

		log.Printf("Hung up session: %s", session.GetId())
	})

	if err := server.Start(); err != nil {
		panic(err)
	}
}
