package cmd

import (
	"github.com/aromanov-pix/aws-mfa-login/action"
	"github.com/spf13/cobra"
)

var clusterCommand = &cobra.Command{
	Use:   "cluster [setup|view]",
	Short: "view or setup your kubeconfig",
	Long:  "view or setup your kubeconfig",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var clusterSetupCommand = &cobra.Command{
	Use:   "setup",
	Short: "setup your kubeconfig",
	Long:  "setup your kubeconfig",
	Run: func(cmd *cobra.Command, args []string) {
		updater := &action.KubeConfigUpdater{}
		updater.Init()
		updater.SetupClusters()
	},
}

var clusterViewCommand = &cobra.Command{
	Use:   "view",
	Short: "view all cluster names and metadata",
	Long:  "view all cluster names and metadata",
	Run: func(cmd *cobra.Command, args []string) {
		updater := &action.KubeConfigUpdater{}
		updater.Init()
		updater.ListClusters()
	},
}

func init() {
	rootCmd.AddCommand(clusterCommand)
	clusterCommand.AddCommand(clusterSetupCommand, clusterViewCommand)
}
