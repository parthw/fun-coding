package cmd

import (
	"fmt"

	"github.com/parthw/kubernetes-endpoints-service/internal/service"
	"github.com/spf13/cobra"
)

var (
	inCluster  bool
	outCluster bool

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "To start service",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if inCluster == outCluster {
				fmt.Println("Please set one flag")
			} else if inCluster {
				service.CountPods(inCluster)
			} else {
				service.CountPods(inCluster)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&inCluster, "in-cluster", "i", false, "To start server inside kubernetes cluster")
	startCmd.Flags().BoolVarP(&outCluster, "out-cluster", "o", false, "To start server outside kubernetes cluster")
}
