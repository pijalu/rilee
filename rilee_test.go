package rilee

import (
	"reflect"
	"testing"
)

type testCase struct {
	input    []int
	expected []int
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
		actual := Encode(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("Expected %v but got %v", c.expected, actual)
		}
	}
}

func TestEncodeLong(t *testing.T) {
	input := make([]int, 1000)
	expected := []int{1000, 0}

	actual := Encode(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
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
		actual, err := Decode(c.input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("Expected %v but got %v", c.expected, actual)
		}
	}
}

func TestDecodeErr(t *testing.T) {
	cases := [][]int{
		[]int{2, 1, 0},
	}

	for _, c := range cases {
		_, err := Decode(c)
		if err == nil {
			t.Fatalf("Expected error but got none !")
		}
	}
}
