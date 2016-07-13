package config

import (
	"strings"
)

var (
	// Database is the database connection config, supplied through environment variables.
	Database databaseConfig
)

type databaseConfig struct {
	Name     string `default:"my_api"`
	Username string
	Password string
	Host     string `default:"localhost"`
	Port     string `default:"5432"`
	SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
}

func (c databaseConfig) String() string {
	params := map[string]string{
		"dbname":   c.Name,
		"user":     c.Username,
		"password": c.Password,
		"host":     c.Host,
		"port":     c.Port,
		"sslmode":  c.SSLMode,
	}

	var parts []string

	for k, v := range params {
		if v != "" {
			parts = append(parts, k+"="+v)
		}
	}

	return strings.Join(parts, " ")
}
