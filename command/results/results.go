package results

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Results struct {
	TotalRequests   int
	SuccessRequests int
	StatusCodes     map[int]int
}

func (r Results) GenerateReport(results Results, duration time.Duration) {	
	fmt.Println()
	red := color.New(color.FgRed)
	red.Println("TEST RESULT")
	fmt.Printf("Test duration: %.2fs\n", duration.Seconds())
	fmt.Printf("Successful Requests: %d\n", results.SuccessRequests)
	fmt.Println()	

	headerFmt := color.New(color.FgGreen).SprintfFunc()
	tbl2 := table.New("Status", "Total")
	tbl2.WithHeaderFormatter(headerFmt)
	for code, count := range results.StatusCodes {
		tbl2.AddRow(code, count)
	}
	tbl2.Print()
	fmt.Println()	
}
