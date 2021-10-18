package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Short: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
	Version: "1.0.0",
}

func main() {
	os.Exit(Execute())
}

func Execute() int {
	if err := cmd.Execute(); err != nil {
		return 1
	}
	return 0
}
