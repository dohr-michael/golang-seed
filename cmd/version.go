package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
)

var (
	BuildVersion  string = "v0.0.0"
	BuildRevision string = "ac7c127b4c24fba18a21601c3838ed0a05d29db9"
	BuildTime     string = "2018-08-22T09:17:08+02:00"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Version:  %s\n", viper.GetString("build.version"))
		fmt.Printf("Revision: %s\n", viper.GetString("build.revision"))
		fmt.Printf("Time:     %s\n", viper.GetString("build.time"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
