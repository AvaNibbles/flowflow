package commands

import (
	"github.com/avanibbles/flowflow/internal/server"
	"github.com/avanibbles/flowflow/pkg/config"
	"github.com/spf13/cobra"
)

var serveCommand = &cobra.Command{Use: "serve", RunE: func(cmd *cobra.Command, args []string) error {
	config, err := config.LoadConfigFromViper()
	if err != nil {
		return err
	}

	srv, err := server.NewServer(config)
	if err != nil {
		return err
	}

	return srv.Run()
}}
