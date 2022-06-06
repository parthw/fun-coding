package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	initCmdFlag := flag.NewFlagSet("init", flag.ExitOnError)
	initCmdHttpPortFlag := initCmdFlag.Int("http-port", 8000, "Port to check health status")
	initListenPortFlag := initCmdFlag.Int("listen-port", 8100, "Port to listen the gossip connection")

	joinCmdFlag := flag.NewFlagSet("join", flag.ExitOnError)
	joinCmdHttpPortFlag := joinCmdFlag.String("http-port", "9000", "Port to check health status")
	joinCmdClusterKeyFlag := joinCmdFlag.String("cluster-key", "", "Cluster key to join the cluster")
	joinCmdKnownIPFlag := joinCmdFlag.String("known-ip-addr", "", "IP and PORT of any live node to join the cluster")
	joinCmdListenPortFlag := joinCmdFlag.String("listen-port", "9100", "Port to listen the gossip connection")

	switch os.Args[1] {

	case "join":
		joinCmdFlag.Parse(os.Args[2:])
		JoinCluster(*joinCmdHttpPortFlag, *joinCmdClusterKeyFlag, *joinCmdKnownIPFlag, *joinCmdListenPortFlag)
	case "init":
		initCmdFlag.Parse(os.Args[2:])
		InitGossipCluster(*initCmdHttpPortFlag, *initListenPortFlag)
	default:
		fmt.Println("expected 'join' or 'init' subcommands")
		os.Exit(1)
	}

	os.Exit(0)

}
