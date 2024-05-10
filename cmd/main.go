package main

import (
	"context"
	"github.com/luongdev/switcher/freeswitch"
	"github.com/luongdev/switcher/freeswitch/interfaces"
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

	_, err := c.Build()

	if err != nil {
		panic(err)
	}

	co := freeswitch.OutboundConfig{
		Host: "10.8.0.2",
		Port: 65022,
	}

	server := co.Build()
	server.OnSessionStarted(func(ctx context.Context, session interfaces.Session) {
		log.Printf("Session started: %s", session.GetId())

		if err := session.Answer(ctx); err != nil {
			log.Printf("Failed to answer session: %s", err)
			return
		}

		log.Printf("Answered session: %s", session.GetId())
	})

	if err := server.Start(); err != nil {
		panic(err)
	}
}
