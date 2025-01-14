package config

import "go.uber.org/dig"

type HlsProvider struct {
	dig.Out

	ServicePort ServicePort
}

type ServicePort string

func (s ServicePort) String() string {
	return string(s)
}
