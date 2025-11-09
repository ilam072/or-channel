package main

import (
	"fmt"
	"github.com/ilam072/or-channel/or"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or.Or(
		sig(2*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
	)
	fmt.Println("done after", time.Since(start))
}
