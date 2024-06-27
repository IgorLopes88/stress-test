package cli

import (
	"os"

	"github.com/igorlopes88/goexpert-stresstest/command/stresstest"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go run main.go",
	Short: "run stress test a given url",
	Long:  "run stress test a given url with a number of requests and concurrency",
	Run: func(cmd *cobra.Command, args []string) {
		test := stresstest.StressTest{
			Url:         urlTest,
			Requests:    requestsTest,
			Concurrency: concurrencyTest,
			Cmd:         cmd,
		}
		test.Execute()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var urlTest string
var requestsTest, concurrencyTest int

func init() {
	rootCmd.Flags().StringVarP(&urlTest, "url", "u", "", "URL of the service to be tested")
	rootCmd.Flags().IntVarP(&requestsTest, "requests", "r", 100, "total number of requests")
	rootCmd.Flags().IntVarP(&concurrencyTest, "concurrency", "c", 10, "number of simultaneous calls")
	rootCmd.MarkFlagRequired("url")
}
