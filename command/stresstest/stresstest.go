package stresstest

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/igorlopes88/goexpert-stresstest/command/hourglass"
	"github.com/igorlopes88/goexpert-stresstest/command/httprequest"
	"github.com/igorlopes88/goexpert-stresstest/command/results"
	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

type StressTest struct {
	Url         string
	Requests    int
	Concurrency int
	Cmd         *cobra.Command
}

func (t *StressTest) start() {
	blue := color.New(color.FgHiBlue)
	fmt.Println()
	blue.Println("-- RUN STRESS TEST -->")
	fmt.Printf("Url: %s\n", t.Url)
	fmt.Printf("Requests: %d\n", t.Requests)
	fmt.Printf("Concurrency: %d\n", t.Concurrency)
	fmt.Println("")
	t.urlValidador(t.Cmd)
}

func (t *StressTest) testing() results.Results {
	result := results.Results{
		TotalRequests:   t.Requests,
		SuccessRequests: 0,
		StatusCodes:     make(map[int]int),
	}
	concurrency := make(chan struct{}, t.Concurrency)

	for i := 0; i < t.Requests; i++ {
		wg.Add(1)
		concurrency <- struct{}{}
		go func() {
			defer wg.Done()
			code, err := httprequest.HttpRequest(t.Url)
			result.StatusCodes[code]++
			if err == nil {
				result.SuccessRequests++
			}
			<-concurrency
		}()
	}
	wg.Wait()
	return result
}

func (t *StressTest) Execute() {
	t.start()
	spin := hourglass.Hourglass{}
	spin.Start()
	result := t.testing()
	spin.Stop()
	result.GenerateReport(result, spin.Duration)
}

func (t *StressTest) urlValidador(cmd *cobra.Command) {
	_, err := url.ParseRequestURI(t.Url)
	if err != nil {
		print("\nError: invalid url for testing\n\n")
		cmd.Help()
		os.Exit(0)
	}
	u, err := url.Parse(t.Url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		print("\nError: invalid url for testing\n\n")
		cmd.Help()
		os.Exit(0)
	}
}
