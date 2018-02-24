package rilee

// Encode encodes a channel of int as RLE to an new channel.
func Encode(input chan int) chan int {
	// Since we are going to write pairs, lets make buffer the channel by 2
	output := make(chan int, 2)
	go func() {
		// When done, close the channel
		defer close(output)
		cnt := 1
		prev, ok := <-input
		if !ok {
			// Let's consider empty as empty
			return
		}

		// Process all channel values
		for current := range input {
			if prev == current {
				cnt++
			} else {
				output <- cnt
				output <- prev
				prev = current
				cnt = 1
			}
		}
		// Push left over
		output <- cnt
		output <- prev
	}()

	return output
}

// Decode decodes a channel of int (RLE) and return as a channel.
// if an error occurs during, the boolean pointer return will be set to true
func Decode(input chan int) (chan int, *bool) {
	output := make(chan int)
	decodeError := false

	go func() {
		// Close the output channel when we are done
		defer close(output)
		for {
			cnt, ok := <-input
			// We have decoded everything
			if !ok {
				return
			}
			value, ok := <-input
			// We miss a value or the count is non-sense
			if !ok || cnt <= 0 {
				decodeError = true
				return
			}
			// Decode !
			for i := 0; i < cnt; i++ {
				output <- value
			}
		}
	}()

	return output, &decodeError
}
