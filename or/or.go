package or

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
