package util

import (
	"github.com/fatih/structs"
	"github.com/jeremywohl/flatten"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func SetupConfigFromEnv(s interface{}) error {
	// Transform config struct to map
	confMap := structs.Map(s)

	// Flatten nested conf map
	flat, err := flatten.Flatten(confMap, "", flatten.DotStyle)
	if err != nil {
		return errors.Wrap(err, "Unable to flatten config")
	}

	// Bind each conf fields to environment vars
	for key := range flat {
		err := viper.BindEnv(key)
		if err != nil {
			return errors.Wrapf(err, "Unable to bind env var: %s", key)
		}
	}

	return nil
}
