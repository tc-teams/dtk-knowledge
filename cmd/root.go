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

}

func config() {

	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
	}

}
