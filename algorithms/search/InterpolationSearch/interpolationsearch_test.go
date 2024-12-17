package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd


 */
func TestinterpolationSearch(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		query    int
		expected int
	}{
		{
			name:     "Scenario 1: Search for an element present in a sorted array",
			array:    []int{1, 3, 5, 7, 9, 11},
			query:    7,
			expected: 3,
		},
		{
			name:     "Scenario 2: Search for an element not present in a sorted array",
			array:    []int{2, 4, 6, 8, 10},
			query:    5,
			expected: -1,
		},
		{
			name:     "Scenario 3: Search in an empty array",
			array:    []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "Scenario 4: Search for the smallest element in a sorted array",
			array:    []int{1, 2, 3, 4, 5},
			query:    1,
			expected: 0,
		},
		{
			name:     "Scenario 5: Search for the largest element in a sorted array",
			array:    []int{3, 4, 5, 6, 7},
			query:    7,
			expected: 4,
		},
		{
			name:     "Scenario 6: Search for an element in a single-element array",
			array:    []int{9},
			query:    9,
			expected: 0,
		},
		{
			name:     "Scenario 7: Search for a non-existent element with equal lower and upper bounds in the array",
			array:    []int{5, 5, 5, 5, 5},
			query:    6,
			expected: -1,
		},
		{
			name:     "Scenario 8: Search for an element in a uniformly distributed array",
			array:    []int{1, 3, 5, 7, 9},
			query:    5,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := interpolationSearch(tt.array, tt.query)
			if result != tt.expected {
				t.Errorf("failed %s: expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("passed %s: expected %d, got %d", tt.name, tt.expected, result)
			}
		})
	}
}

