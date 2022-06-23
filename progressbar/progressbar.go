package progressbar

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"math"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	clearLine  = "\033[1A\033\r"
	deleteLine = "\033[M"

	// states:
	startup   = 0
	running   = 1
	complete  = 2
	cancelled = 3
	shutdown  = 4
)

type progressbar struct {
	mu               *sync.Mutex
	name             string
	totLength        int
	progress         int
	startTime        time.Time
	status           int32
	barLength        int
	totalDurationEst time.Duration
	totalTime        time.Duration
}

func newProgressBar(batchName string, length int) *progressbar {
	w := getTerminalWidth()
	w = int(math.Max(5, float64(w-50)))

	return &progressbar{
		mu:        &sync.Mutex{},
		name:      batchName,
		totLength: length,
		progress:  0,
		status:    0,
		barLength: w,
	}
}

// start should be called immediately prior to execution of tasks.
func (b *progressbar) start() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.startTime = time.Now()
	b.setStatus(running)

	go func() {
		for b.readStatus() != shutdown {
			time.Sleep(100 * time.Millisecond)
			b.mu.Lock()
			b.printBar()
			b.mu.Unlock()
		}
	}()

	fmt.Println()

	b.printBar()
}

// increment should be called each time a task is completed.
func (b *progressbar) increment() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.progress++

	// update duration estimate
	completeRatio := float64(b.progress) / float64(b.totLength)
	durRemainingS := time.Now().Sub(b.startTime).Seconds() / completeRatio
	totDuration, err := time.ParseDuration(fmt.Sprintf("%fs", durRemainingS))
	if err != nil {
		panic(err)
	}
	b.totalDurationEst = totDuration

	if b.progress == b.totLength && b.readStatus() == running {
		b.setStatus(complete)
		b.totalTime = time.Now().Sub(b.startTime)
		b.printBar()
	}
}

// cleanup should be called in a defer statement.
func (b *progressbar) cleanup() {
	time.Sleep(100 * time.Millisecond)
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.readStatus() == complete || b.readStatus() == cancelled || b.readStatus() == shutdown {
		return
	}

	b.setStatus(cancelled)
	b.totalTime = time.Now().Sub(b.startTime)
	b.printBar()
}

// cancel should be called when exiting unexpectedly.
func (b *progressbar) cancel(err error) {
	//time.Sleep(100 * time.Millisecond)
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.readStatus() == complete || b.readStatus() == cancelled || b.readStatus() == shutdown {
		return
	}

	b.setStatus(cancelled)
	b.totalTime = time.Now().Sub(b.startTime)
	b.printBar()
	fmt.Printf("%s: %s\n", b.name, err.Error())
}

// printBar: THREAD-UNSAFE!
func (b *progressbar) printBar() {

	if b.readStatus() == shutdown {
		return
	}

	percentageComplete := (float64(b.progress) / float64(b.totLength)) * 100

	// build bar string
	bar0 := fmt.Sprintf("%s: %d / %d", b.name, b.progress, b.totLength)

	bar2 := strings.TrimSpace(fmt.Sprintf("%3.0f", math.Floor(percentageComplete)))
	bar2 += "%"
	switch b.readStatus() {
	case startup:
		bar2 += " - Remaining: ??"

	case running:
		if percentageComplete == 0 {
			bar2 += " - Remaining: ??"
		} else {
			bar2 += " - Remaining: "
			bar2 += formatDuration(b.totalDurationEst - time.Now().Sub(b.startTime))
		}

	case complete:
		bar2 += " - Completed: "
		bar2 += formatDuration(b.totalTime)
		b.setStatus(shutdown)

	case cancelled:
		bar2 += " - Cancelled: "
		bar2 += formatDuration(b.totalTime)
		b.setStatus(shutdown)

	case shutdown:
		fmt.Println()
		return
	}

	barLength := getTerminalWidth() - len(bar0) - len(bar2) - 5
	bar1 := " - "

	if barLength > 4 {
		bar1 = " ["
		d := 100 / float64(barLength)
		for i := 0; i < barLength; i++ {
			c := " "
			if i < int(percentageComplete/d) {
				c = "="
			} else if i == int(percentageComplete/d) && i == barLength-1 {
				c = "="
			} else if i == int(percentageComplete/d) {
				c = ">"
			}
			bar1 += c
		}
		bar1 += "] "
	}

	fmt.Print(clearLine + deleteLine)
	fmt.Print(bar0 + bar1 + bar2 + "\n")
}

func (b *progressbar) isRunning() bool {
	return atomic.LoadInt32(&b.status) == running
}

func (b *progressbar) setStatus(status int32) {
	atomic.StoreInt32(&b.status, status)
}

func (b *progressbar) readStatus() int32 {
	return atomic.LoadInt32(&b.status)
}

// UTIL

func formatDuration(dur time.Duration) string {
	s := truncateToSignificantDigits(dur.Seconds(), 3)
	d, _ := time.ParseDuration(fmt.Sprintf("%fs", s))

	ds := fmt.Sprintf("%s", d)

	if len(ds) < 8 {
		ds = fmt.Sprintf("%-8s", ds)
	}

	return ds
}

func truncateToSignificantDigits(input float64, n uint) float64 {
	check := math.Pow10(int(n) - 1)
	sign := 1.0

	if input == 0 {
		return 0
	} else if input < 0 {
		sign = -1.0
	}

	i := 0
	m := sign * math.Pow10(i) // first iteration m = 1

	for input*m < check {
		m = sign * math.Pow10(i)
		i++
	}

	return math.Trunc(input*m) / m
}

func getTerminalWidth() int {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, _, err = terminal.GetSize(0)
		if err != nil {
			return 20
		}
	}

	if width > 100 {
		return 100
	}

	return width
}
