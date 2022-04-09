package flowflow

import (
	"os"
	"strings"

	"github.com/avanibbles/flowflow/internal/util"
	"github.com/avanibbles/flowflow/pkg/config"

	"github.com/pkg/errors"

	"github.com/avanibbles/flowflow/internal"
	"github.com/avanibbles/flowflow/internal/flowflow/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/joho/godotenv/autoload"
)

var (
	cfgFile string

	Cmd = &cobra.Command{
		Use:     "flowflow",
		Short:   "flow your model flow with flowflow",
		Version: internal.BuildVersion(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		}}
)

func init() {
	Cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flowflow.yaml)")

	commands.SetupCommand(Cmd)
}

func initConfig() error {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".flowflow.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return err
		}
	}

	viper.SetEnvPrefix("FLOWFLOW")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	cfg := config.NewConfig()
	if err := util.SetupConfigFromEnv(cfg); err != nil {
		panic(err)
	}

	return nil
}
