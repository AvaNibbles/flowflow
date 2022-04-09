package config

import (
	"os"
	"path/filepath"
)

func DefaultConfigDirectory() string {
	if home, err := os.UserHomeDir(); err != nil {
		confDir, err := os.Getwd()
		if err != nil {
			// todo: why would Getwd reurn errors?
			// are there ways to further handle this?
			// in the use cases where we would want  this function, we wouldn't have much more we could
			// do to resolve any kind of config dir so just crash
			panic(err)
		}
		return confDir
	} else {
		return home
	}
}

func DefaultConfigFile() string {
	return filepath.Join(DefaultConfigDirectory(), ".flowflow.yaml")
}
