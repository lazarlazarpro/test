package main

import (
	"log"
	"reflect"
	"sort"
	"testing"
)

func setup() {
	err := LoadConfig("../config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(config.PackSizes)))
}

func TestGetPackSizes(t *testing.T) {
	setup()

	tests := []struct {
		input  int
		output map[int]int
	}{
		{
			input: 1,
			output: map[int]int{
				250: 1,
			},
		},
		{
			input: 250,
			output: map[int]int{
				250: 1,
			},
		},
		{
			input: 251,
			output: map[int]int{
				500: 1,
			},
		},
		{
			input: 501,
			output: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			input: 12001,
			output: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}

	for _, test := range tests {
		result := getPackSizes(test.input)
		if !reflect.DeepEqual(*result, test.output) {
			t.Errorf("For input %d, expected %+v, but got %+v", test.input, test.output, *result)
		}
	}
}
