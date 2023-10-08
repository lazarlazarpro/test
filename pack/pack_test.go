package pack

import (
	"reflect"
	"testing"
)

func TestGetPackSizes(t *testing.T) {

	packSizes := []int{250, 500, 1000, 2000, 5000}

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

	p := NewPacker(packSizes)
	for _, test := range tests {
		result := p.GetPackSizes(test.input)
		if !reflect.DeepEqual(*result, test.output) {
			t.Errorf("For input %d, expected %+v, but got %+v", test.input, test.output, *result)
		}
	}
}
