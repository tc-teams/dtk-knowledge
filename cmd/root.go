package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	configFileFlag string
)

var rootCmd = &cobra.Command{
	Use:   "Finder",
	Short: "A brief description of your application",
	Long:  "A longer description: Web Application for track news related by covid",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(config)
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.Flags().StringVarP(&configFileFlag, "config", "f", "./config.yaml", "The path to the config file to use.")
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

}

func config() {

	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
	}

}
