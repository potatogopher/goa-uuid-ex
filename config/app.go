package config

import (
	"errors"
	"time"
)

var (
	// App is the general app config, supplied through environment variables.
	App appConfig

	validEnvironments = map[string]bool{
		"test":        true,
		"development": true,
		"staging":     true,
		"production":  true,
	}
)

type appConfig struct {
	Environment     string        `default:"development"`
	StatsDHost      string        `envconfig:"STATSD_HOST"`
	MetricsInterval time.Duration `envconfig:"METRICS_INTERVAL" default:"10s"`
}

func (c appConfig) Validate() error {
	if _, ok := validEnvironments[App.Environment]; !ok {
		return errors.New("not a valid environment")
	}
	return nil
}
