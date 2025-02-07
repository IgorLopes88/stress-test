package stresstest

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
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
	Begin       time.Time
	Final       time.Time
	Duration    time.Duration
}

func (t *StressTest) start() {
	blue := color.New(color.FgHiBlue)
	t.Begin = time.Now()
	fmt.Println()
	blue.Println("-- RUN STRESS TEST -->")
	fmt.Printf("Url: %s\n", t.Url)
	fmt.Printf("Requests: %d\n", t.Requests)
	fmt.Printf("Concurrency: %d\n", t.Concurrency)
	fmt.Println("")
	t.urlValidador(t.Cmd)
	fmt.Println("Testing...")
}

func (t *StressTest) testing() results.Results {
	m := sync.Mutex{}
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
			m.Lock()
			result.StatusCodes[code]++
			if err == nil {
				result.SuccessRequests++
			}
			m.Unlock()
			<-concurrency
		}()
	}
	wg.Wait()
	t.Final = time.Now()
	t.Duration = t.Final.Sub(t.Begin)
	return result
}

func (t *StressTest) Execute() {
	t.start()
	result := t.testing()
	result.GenerateReport(result, t.Duration)
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
