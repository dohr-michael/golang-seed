package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var startCmd = &cobra.Command{
	Use: "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Start :)")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
