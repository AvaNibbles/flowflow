package commands

import (
	"fmt"
	"os"

	"github.com/avanibbles/flowflow/pkg"
	"github.com/avanibbles/flowflow/pkg/config"

	"github.com/avanibbles/flowflow/internal"
	"github.com/spf13/cobra"
)

func SetupCommand(cmd *cobra.Command) {
	cmd.AddCommand(versionCommand, serveCommand, configCommand)
}

var (
	versionCommand = &cobra.Command{
		Use:   "version",
		Short: "Print the current flowflow version",
		RunE: func(_ *cobra.Command, _ []string) error {
			conf, err := config.LoadConfigFromViper()
			cobra.CheckErr(err)

			var serverVersion string
			if len(conf.Client.Host) > 0 {

				api, err := pkg.NewApiClient(conf)
				cobra.CheckErr(err)

				resp, err := api.Version.GetAPIV1Version(nil)

				if err != nil {
					serverVersion = fmt.Sprintf("server: err: %s", err.Error())
				} else {
					serverVersion = fmt.Sprintf("%s-%s (%s)", resp.Payload.Version, resp.Payload.CommitHash, resp.Payload.BuildTimestamp)
					serverVersion = fmt.Sprintf("server: %s", serverVersion)
				}
			}

			localVersion := fmt.Sprintf("local: %s", internal.BuildVersion())

			_, err = fmt.Fprintln(os.Stdout, localVersion)
			if err != nil {
				return err
			}

			if len(serverVersion) != 0 {
				_, err = fmt.Fprintln(os.Stdout, serverVersion)
				return err
			}

			return nil
		}}
)
