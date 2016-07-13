package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func attemptLoad(filenames ...string) error {
	for _, filename := range filenames {
		err := godotenv.Load(filename)
		if err != nil && os.IsExist(err) {
			return err
		}
	}

	return nil
}

// LoadConfig loads in "./.env" into local environment,
// before parsing environment varables.
func LoadConfig() error {
	if err := attemptLoad("../.env", "./.env"); err != nil {
		return err
	}

	err := envconfig.Process("thing", &App)
	if err != nil {
		return err
	}

	if err := App.Validate(); err != nil {
		return err
	}

	err = envconfig.Process("my_database", &Database)
	if err != nil {
		return err
	}

	return nil
}
