package hourglass

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type Hourglass struct {
	Begin    time.Time
	Final    time.Time
	Duration time.Duration
}

var Spin = spinner.New([]string{".", "..", "..."}, 500*time.Millisecond)

func (h *Hourglass) Start() {
	h.Begin = time.Now()
	Spin.Prefix = "Testing"
	green := color.New(color.FgHiGreen)
	Spin.FinalMSG = green.Sprintln("Test finished!")
	Spin.Start()
}

func (h *Hourglass) Stop() {
	Spin.Stop()
	h.Final = time.Now()
	h.Duration = h.Final.Sub(h.Begin)
}
