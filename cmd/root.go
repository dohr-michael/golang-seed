package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
	"fmt"
	"log"
	"path/filepath"
	"os"
	"io/ioutil"
)

var configFile string
var verbose bool

// TODO Change me
const cmdName = "golang-seed"

var rootCmd = &cobra.Command{
	Use:   cmdName,
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", fmt.Sprintf("config file (default \"$HOME/.%s.yml\")", cmdName))
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Set Default Viper configs.
	viper.SetDefault("build.version", BuildVersion)
	viper.SetDefault("build.revision", BuildRevision)
	viper.SetDefault("build.time", BuildTime)
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}
		filename := filepath.Join(home, fmt.Sprintf(".%s.yml", cmdName))
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			// Default config file.
			configYml := `
`
			err = ioutil.WriteFile(filename, []byte(configYml), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}

		viper.SetConfigName(fmt.Sprintf(".%s", cmdName))
		viper.AddConfigPath(home)
	}
	viper.SetEnvPrefix(cmdName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
