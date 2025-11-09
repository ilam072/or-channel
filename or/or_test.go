package or

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-Or(
		sig(2*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
	)

	elapsed := time.Since(start)
	if elapsed >= 2*time.Second {
		t.Errorf("Or took too long: %v", elapsed)
	}
}

func TestOrSingleChannel(t *testing.T) {
	c := make(chan interface{})
	go func() { close(c) }()

	select {
	case <-Or(c):
	case <-time.After(time.Second):
		t.Error("Or did not close the single channel")
	}
}

func TestOrNoChannels(t *testing.T) {
	select {
	case <-Or():
	case <-time.After(time.Second):
		t.Error("Or did not close the empty channel list")
	}
}
