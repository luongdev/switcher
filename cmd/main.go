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
	server.OnSessionStarted(func(ctx context.Context, client interfaces.Client, event interfaces.Event) {
		log.Printf("Session opened")
	})

	if err := server.Start(); err != nil {
		panic(err)
	}
}
