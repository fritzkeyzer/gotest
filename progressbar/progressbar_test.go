package progressbar

import (
	"testing"
	"time"
)

func TestProgressBar(t *testing.T) {
	pb := newProgressBar("hello", 100)
	pb.start()
	defer pb.cleanup()

	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		pb.increment()
	}

}
