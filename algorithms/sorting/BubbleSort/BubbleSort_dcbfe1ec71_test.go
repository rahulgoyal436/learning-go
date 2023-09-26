package BubbleSort

import (
	"reflect"
	"testing"
)

func TestBubbleSort_dcbfe1ec71(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "Test case 1: Normal case",
			input:          []int{5, 3, 2, 8, 6, 7},
			expectedOutput: []int{2, 3, 5, 6, 7, 8},
		},
		{
			name:           "Test case 2: Edge case with empty array",
			input:          []int{},
			expectedOutput: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := bubbleSort(tc.input)
			if !reflect.DeepEqual(output, tc.expectedOutput) {
				t.Logf("For input: %v", tc.input)
				t.Errorf("Expected: %v, but got: %v", tc.expectedOutput, output)
			} else {
				t.Logf("Success for input: %v", tc.input)
			}
		})
	}
}
