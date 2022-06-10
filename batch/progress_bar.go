package batch

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"math"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

const (
	clearLine     = "\033[1A\033\r"
	deleteLine    = "\033[M"
	colorClear    = "\033[00;32m"
	colorGreen    = "\033[01;32m"
	backSpace     = "\b"
	cursorSave    = "\033[7"
	cursorRestore = "\033[8"
	deleteToEnd   = "\033[0J"

	// states:
	startup   = 0
	running   = 1
	complete  = 2
	cancelled = 3
	shutdown  = 4
)

type ProgressBar struct {
	mu               *sync.Mutex
	totLength        int
	progress         int
	startTime        time.Time
	status           int32
	barLength        int
	lastPrintLength  int
	totalDurationEst time.Duration
	totalTime        time.Duration
}

func NewProgressBar(length int) *ProgressBar {
	w := getWidth()
	w = int(math.Max(5, float64(w-50)))

	return &ProgressBar{
		mu:        &sync.Mutex{},
		totLength: length,
		progress:  0,
		status:    0,
		barLength: w,
	}
}

// Start should be called immediately prior to execution of tasks.
func (b *ProgressBar) Start() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.startTime = time.Now()
	atomic.StoreInt32(&b.status, running)

	go func() {
		for b.readStatus() != shutdown {
			time.Sleep(100 * time.Millisecond)
			b.printBar()
		}
	}()

	fmt.Println()

	b.printBar()
}

// Increment should be called each time a task is completed.
func (b *ProgressBar) Increment() {
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
		atomic.StoreInt32(&b.status, complete)
		b.totalTime = time.Now().Sub(b.startTime)
		b.printBar()
	}
}

// Cancel should be called when exiting unexpectedly.
func (b *ProgressBar) Cancel() {
	time.Sleep(100 * time.Millisecond)
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.readStatus() == complete || b.readStatus() == cancelled || b.readStatus() == shutdown {
		return
	}

	atomic.StoreInt32(&b.status, cancelled)
	b.totalTime = time.Now().Sub(b.startTime)
	b.printBar()
}

// printBar: THREAD-UNSAFE!
func (b *ProgressBar) printBar() {

	if b.readStatus() == shutdown {
		return
	}

	percentageComplete := (float64(b.progress) / float64(b.totLength)) * 100

	// build bar string
	bar := clearLine + deleteLine + fmt.Sprintf("%d / %d", b.progress, b.totLength)

	bar += " ["
	d := 100 / float64(b.barLength)
	for i := 0; i < b.barLength; i++ {
		c := " "
		if i < int(percentageComplete/d) {
			c = "="
		} else if i == int(percentageComplete/d) {
			//c = string(rune(187))
			c = ">"
		}
		bar += c
	}
	bar += "] "
	bar += fmt.Sprintf("%3.0f", math.Floor(percentageComplete))
	bar += "%"

	switch b.readStatus() {
	case startup:
		bar += " - Remaining: ??"

	case running:
		if percentageComplete == 0 {
			bar += " - Remaining: ??"
		} else {
			bar += " - Remaining: "
			bar += formatDuration(b.totalDurationEst - time.Now().Sub(b.startTime))
		}

	case complete:
		bar += " - Completed: "
		bar += formatDuration(b.totalTime)
		atomic.StoreInt32(&b.status, shutdown)

	case cancelled:
		bar += " - Cancelled: "
		bar += formatDuration(b.totalTime)
		atomic.StoreInt32(&b.status, shutdown)

	case shutdown:
		return
	}

	bar += "\n"

	b.lastPrintLength = len(bar)
	fmt.Print(bar)
}

// isRunning: THREAD-UNSAFE!
func (b *ProgressBar) isRunning() bool {
	return atomic.LoadInt32(&b.status) == running
}

func (b *ProgressBar) readStatus() int32 {
	return atomic.LoadInt32(&b.status)
}

func formatDuration(dur time.Duration) string {
	//var d time.Duration
	//if s > 100 {
	//	d = dur.Truncate(1000 * time.Millisecond)
	//} else if s > 10 {
	//	d = dur.Truncate(100 * time.Millisecond)
	//} else if s > 1 {
	//	d = dur.Truncate(10 * time.Millisecond)
	//} else if s > 0.1 {
	//	d = dur.Truncate(1 * time.Millisecond)
	//}

	s := sf(dur.Seconds(), 3)
	d, _ := time.ParseDuration(fmt.Sprintf("%fs", s))

	return fmt.Sprintf("%s", d)
}

func sf(input float64, n uint) float64 {
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

func getWidth() int {
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, _, err = terminal.GetSize(int(os.Stderr.Fd()))
		if err != nil {
			width = 20
		}
	}

	return width
}
