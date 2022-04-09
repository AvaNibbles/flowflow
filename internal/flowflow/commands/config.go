package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/avanibbles/flowflow/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	configCommand    = &cobra.Command{Use: "config", Short: "config management commands"}
	configGetCommand = &cobra.Command{
		Use:   "get",
		Short: "get the current configuration",
		Run: func(cmd *cobra.Command, args []string) {
			conf, err := config.LoadConfigFromViper()
			cobra.CheckErr(err)

			b, err := yaml.Marshal(conf)
			cobra.CheckErr(err)

			fmt.Print(string(b))
		}}

	configSetCommand = &cobra.Command{
		Use:   "set key=value...",
		Short: "set one or more values in the current configuration and save it",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, directive := range args {
				split := strings.Split(directive, "=")
				if len(split) != 2 {
					fmt.Printf("directive \"%s\" unprocessable", directive)
					os.Exit(-1)
				}

				viper.Set(split[0], split[1])
			}

			configUsed := viper.ConfigFileUsed()
			if len(configUsed) == 0 {
				configUsed = config.DefaultConfigFile()
				viper.SetConfigFile(configUsed)
			}

			cobra.CheckErr(viper.WriteConfig())
		},
	}
)

func init() {
	configCommand.AddCommand(configGetCommand, configSetCommand)
}
