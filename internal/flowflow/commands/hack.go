package commands

import (
	"fmt"

	"github.com/avanibbles/flowflow/pkg"
	"github.com/avanibbles/flowflow/pkg/client/hack"
	"github.com/avanibbles/flowflow/pkg/config"
	"github.com/spf13/cobra"
)

var (
	errorCode int

	hackCommand   = &cobra.Command{Use: "hack", Short: "dev testing"}
	errorCodeTest = &cobra.Command{
		Use:   "errors",
		Short: "test errors",

		Run: func(cmd *cobra.Command, args []string) {
			conf, err := config.LoadConfigFromViper()
			cobra.CheckErr(err)

			api, err := pkg.NewApiClient(conf)
			cobra.CheckErr(err)

			req := hack.NewGetAPIV1HackErrCodeParams()
			req.SetCode(int64(errorCode))
			resp, err := api.Hack.GetAPIV1HackErrCode(req)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp)
			}
		}}
)

func init() {
	errorCodeTest.Flags().IntVarP(&errorCode, "error", "e", 200, "the error code to produce")

	hackCommand.AddCommand(errorCodeTest)
}
