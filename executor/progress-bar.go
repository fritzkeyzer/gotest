package executor

import (
	"fmt"
	"sync"
	"time"
)

const ClearLine = "\033[1A\033\r" // \033[1A\033[K    --\033[2K\r

type ProgressBar struct {
	mu               *sync.Mutex
	totLength        int
	progress         int
	startTime        time.Time
	status           int
	barLength        int
	totalDurationEst time.Duration
}

func NewBar(length int) *ProgressBar {
	return &ProgressBar{
		mu:        &sync.Mutex{},
		totLength: length,
		progress:  0,
		status:    0,
		barLength: 20,
	}
}

func (b *ProgressBar) Start() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.status == 0 {
		b.startTime = time.Now()
		b.status = 1
	}

	go func() {
		for b.isRunning() {
			time.Sleep(1 * time.Second)
			b.Update()
		}
	}()

	b.printBar()
}

func (b *ProgressBar) Increment() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.status == 0 {
		b.startTime = time.Now()
		b.status = 1
	}
	b.progress++

	// update duration estimate
	completeRatio := float64(b.progress) / float64(b.totLength)
	durRemainingS := time.Now().Sub(b.startTime).Seconds() / completeRatio
	totDuration, err := time.ParseDuration(fmt.Sprintf("%.0fs", durRemainingS))
	if err != nil {
		panic(err)
	}
	b.totalDurationEst = totDuration

	b.printBar()
}

func (b *ProgressBar) Update() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.printBar()
}

func (b *ProgressBar) Cancel() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.status != 1 {
		return
	}
	b.status = 2
}

func (b *ProgressBar) printBar() {
	percentageComplete := (float64(b.progress) / float64(b.totLength)) * 100

	// print ProgressBar
	bar := ClearLine + fmt.Sprintf("%d / %d", b.progress, b.totLength)

	bar += " ["
	d := 100 / float64(b.barLength)
	for i := 0; i < b.barLength; i++ {
		c := " "
		if i < int(percentageComplete/d) {
			c = "="
		} else if i == int(percentageComplete/d) {
			c = ">"
		}
		bar += c
	}
	bar += "] "
	bar += fmt.Sprintf("%3.0f", percentageComplete)
	bar += "%"

	if percentageComplete == 0 {
		bar += " - Remaining: ??"
	} else if percentageComplete == 100 {
		comDurS := time.Now().Sub(b.startTime).Seconds()
		comDur, err := time.ParseDuration(fmt.Sprintf("%.0fs", comDurS))
		if err != nil {
			panic(err)
		}
		bar += fmt.Sprintf(" - Completed: %v", comDur)
	} else {
		currDurS := time.Now().Sub(b.startTime).Seconds()
		durRemainingS := b.totalDurationEst.Seconds() - currDurS
		durRemaining, err := time.ParseDuration(fmt.Sprintf("%.0fs", durRemainingS))
		if err != nil {
			panic(err)
		}
		bar += " - Remaining: "
		bar += fmt.Sprint(durRemaining)
	}

	fmt.Println(bar)
}

func (b *ProgressBar) isRunning() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.status == 1
}
