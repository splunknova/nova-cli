package cmd

import (
	"github.com/spf13/cobra"
	"github.com/splunknova/nova-cli/src"
	log "github.com/Sirupsen/logrus"
	"os"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available metrics",
	Run: func(cmd *cobra.Command, args []string) {
		clientID, clientSecret, err := src.GetCredentials(NovaURL)
		if err != nil {
			log.Error(err)
			log.Infof("Please run `nova login`")
			os.Exit(1)
		}
		authHeader := src.GetBasicAuthHeader(clientID, clientSecret)

		m := src.NewNovaMetricsSearch(NovaURL, authHeader)

		data, err := m.GetLs()
		if err != nil {
			os.Exit(1)
		}
		if table, _ := rootCmd.Flags().GetBool("table"); table {
			data.PrintTable()
		} else {
			data.PrintList()
		}
	},
}

func init() {
	metricCmd.AddCommand(lsCmd)
}
