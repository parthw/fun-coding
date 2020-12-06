/*
Package cmd Copyright © 2020 Parth

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"

	"example.com/grpcFinale/internal/client"
	"example.com/grpcFinale/internal/server"
	"github.com/spf13/cobra"
)

var (
	// Used Flags
	startServer bool
	startClient bool

	// startCmd represents the start command
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "To start gRPC server or client",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			if startServer == startClient {
				log.Fatalln("Exact one flag required")
			}
			if startServer {
				// start server
				log.Println("Starting gRPC server")
				server.StartServer()
			}
			if startClient {
				// start client
				log.Println("Starting gRPC client")
				client.StartClient()
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().BoolVarP(&startServer, "server", "s", false, "To start server")
	startCmd.Flags().BoolVarP(&startClient, "client", "c", false, "To start client")
}
