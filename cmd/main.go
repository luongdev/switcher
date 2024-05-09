package main

import (
	"context"
	"github.com/luongdev/switcher/freeswitch"
	"github.com/luongdev/switcher/freeswitch/pkg"
	"time"
)

func main() {
	c := freeswitch.InboundConfig{
		Host:           "10.8.0.1",
		Port:           65021,
		Password:       "Simplefs!!",
		ConnectTimeout: time.Millisecond * 2,
	}

	client := pkg.NewClient(c)
	err := client.Connect(context.Background())

	if err != nil {
		panic(err)
	}
}
