package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "application version",
	Long:  `Long description`,
	RunE: func(cmd *cobra.Command, args []string) error {

		file, err := os.Open("./VERSION")
		if err != nil {
			return err
		}
		defer file.Close()

		version, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		fmt.Println("version:", string(version))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
