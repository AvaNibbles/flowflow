package main

import (
	"fmt"
	"os"

	"github.com/avanibbles/flowflow/internal/flowflow"
)

func main() {
	if err := flowflow.Cmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
