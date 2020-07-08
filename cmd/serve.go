package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tc-teams/fakefinder-crawler/app"
)

var	serveCmd = &cobra.Command{
		Use:   "server",
		Short: "serve the api",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := app.Run()

			if err := app.Serve(); err != nil{
				return err
			}
			return nil
		},
	}

func init() {
	rootCmd.AddCommand(serveCmd)
}


