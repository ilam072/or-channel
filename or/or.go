// Package or provides a utility to combine multiple done channels.
// The Or function returns a channel that closes as soon as any of the
// provided channels is closed. This is useful for merging cancellation
// signals from multiple sources.
//
// Example usage:
//
//	done := or.Or(ch1, ch2, ch3)
//	<-done // channel closes when any of ch1, ch2, or ch3 closes
package or

// Or combines multiple done channels into a single channel.
// The returned channel will be closed as soon as any of the input channels close.
func Or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		ch := make(chan interface{})
		close(ch)
		return ch
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			mid := len(channels) / 2
			select {
			case <-Or(channels[:mid]...):
			case <-Or(channels[mid:]...):
			}
		}
	}()
	return orDone
}
