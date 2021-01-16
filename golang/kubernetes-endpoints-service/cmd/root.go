package cmd

import (
	"fmt"
	"os"

	"github.com/parthw/kubernetes-endpoints-service/internal/logger"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubernetes-endpoints-service",
	Short: "To fetch pods endpoints details and perform tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Application to fetch pods endpoints details and perform tasks")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubernetes-endpoints-service.yaml)")

}

// initConfig reads in config file
func initConfig() {

	defaultConfig()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kubernetes-endpoints-service" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kubernetes-endpoints-service")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// To initialize logger
	logger.InitializeLogger()
}

func defaultConfig() {
	viper.SetDefault("env", "dev")
	viper.SetDefault("log.file", "app.log")
	viper.SetDefault("log.file.maxsize", "10") //megabytes
	viper.SetDefault("log.file.maxbackups", "7")
	viper.SetDefault("log.file.maxage", "7") //days
	viper.SetDefault("log.level", "info")
}
