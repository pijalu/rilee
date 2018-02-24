package rilee

import (
	"log"
	"reflect"
	"testing"
)

type testCase struct {
	input    []int
	expected []int
}

func toArray(input chan int) []int {
	result := make([]int, 0, 10)
	for i := range input {
		result = append(result, i)
	}
	return result
}

func toChannel(input []int) chan int {
	result := make(chan int, len(input))
	defer close(result)
	for _, i := range input {
		result <- i
	}
	return result
}

func TestEncode(t *testing.T) {
	cases := []testCase{
		{
			input:    []int{0, 0, 0, 1, 2, 2},
			expected: []int{3, 0, 1, 1, 2, 2},
		}, {
			input:    []int{0, 1, 1, 2, 2, 2},
			expected: []int{1, 0, 2, 1, 3, 2},
		},
	}

	for _, c := range cases {
		actual := toArray(
			Encode(
				toChannel(
					c.input)))
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("Expected %v but got %v", c.expected, actual)
		}
		log.Printf("%v => %v", c.input, actual)
	}
}

func TestEncodeLong(t *testing.T) {
	input := make([]int, 1000)
	expected := []int{1000, 0}

	actual := toArray(
		Encode(
			toChannel(input)))
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
	//log.Printf("%v => %v", input, actual)
}

func TestDecode(t *testing.T) {
	cases := []testCase{
		{
			input:    []int{3, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 0, 1, 2, 2},
		}, {
			input:    []int{1, 0, 2, 1, 3, 2},
			expected: []int{0, 1, 1, 2, 2, 2},
		}, {
			input:    []int{1, 0, 1, 0},
			expected: []int{0, 0},
		},
	}

	for _, c := range cases {
		actualChan, decodeError := Decode(toChannel(c.input))
		actual := toArray(actualChan)
		if *decodeError {
			t.Fatalf("Error during decode")
		}
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("Expected %v but got %v", c.expected, actual)
		}
		log.Printf("%v => %v", c.input, actual)
	}
}

func TestDecodeErr(t *testing.T) {
	cases := [][]int{
		[]int{2, 1, 0},
	}

	for _, c := range cases {
		actualChan, decodeError := Decode(toChannel(c))
		_ = toArray(actualChan)
		if !*decodeError {
			t.Fatalf("Expected error but got none !")
		}
	}
}

// Benchmark a encode
func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toArray(
			Encode(
				toChannel(
					[]int{3, 0, 2, 1, 1, 2})))

	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		actualChan, decodeError := Decode(
			toChannel(
				[]int{3, 0, 1, 1, 2, 2}))
		_ = toArray(actualChan)
		if *decodeError {
			b.FailNow()
		}
	}
}
