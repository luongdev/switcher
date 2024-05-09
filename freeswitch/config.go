package freeswitch

import "time"

type InboundConfig struct {
	Host           string        `yaml:"host"`
	Port           uint16        `yaml:"port"`
	Password       string        `yaml:"password"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
}

type OutboundConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type Config struct {
	Inbound  InboundConfig  `yaml:"inbound"`
	Outbound OutboundConfig `yaml:"outbound"`
}
