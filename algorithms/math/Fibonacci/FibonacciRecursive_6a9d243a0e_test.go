// Test generated by RoostGPT for test go-unit-algo-string using AI Type Open AI and AI Model gpt-4

package Fibonacci

import (
	"testing"
)

func TestFibonacciRecursive_6a9d243a0e(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{-1, -1}, // TODO: Change this to the expected output for negative numbers
	}

	for _, testCase := range testCases {
		result := FibonacciRecursive(testCase.n)
		if result != testCase.expected {
			t.Errorf("FibonacciRecursive(%d) = %d; want %d", testCase.n, result, testCase.expected)
		} else {
			t.Logf("Success: FibonacciRecursive(%d) = %d", testCase.n, result)
		}
	}
}
