package rilee

import "fmt"

func Encode(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}

	output := make([]int, 0, len(input))
	cnt := 1
	var prev int = input[0]

	for _, i := range input[1:] {
		if prev == i {
			cnt++
		} else {
			output = append(output, cnt, prev)
			prev = i
			cnt = 1
		}
	}
	output = append(output, cnt, prev)

	return output
}

func Decode(input []int) ([]int, error) {
	if len(input) == 0 {
		return []int{}, nil
	}

	if len(input)%2 != 0 {
		return nil, fmt.Errorf("Invalid RLE length: %d", len(input))
	}

	output := make([]int, 0, len(input))
	for i := 0; i < len(input); i += 2 {
		if input[i] <= 0 {
			return nil, fmt.Errorf("Invalid RLE counter: %d", input[i])
		}
		for j := 0; j < input[i]; j++ {
			output = append(output, input[i+1])
		}
	}

	return output, nil
}
