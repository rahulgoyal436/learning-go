package BinarySearch

import "testing"

/*
ROOST_METHOD_HASH=binarySearch_5149337a0e
ROOST_METHOD_SIG_HASH=binarySearch_7d22ad2576


 */
func TestbinarySearch(t *testing.T) {
	tests := []struct {
		name          string
		arr           []int
		query         int
		expectedIndex int
	}{
		{
			name:          "Element present in the middle",
			arr:           []int{1, 3, 5, 7, 9},
			query:         5,
			expectedIndex: 2,
		},
		{
			name:          "Search for the smallest element",
			arr:           []int{1, 3, 5, 7, 9},
			query:         1,
			expectedIndex: 0,
		},
		{
			name:          "Search for the largest element",
			arr:           []int{1, 3, 5, 7, 9},
			query:         9,
			expectedIndex: 4,
		},
		{
			name:          "Element not present in the array",
			arr:           []int{1, 3, 5, 7, 9},
			query:         4,
			expectedIndex: -1,
		},
		{
			name:          "Empty array",
			arr:           []int{},
			query:         3,
			expectedIndex: -1,
		},
		{
			name:          "Array with one element (element present)",
			arr:           []int{1},
			query:         1,
			expectedIndex: 0,
		},
		{
			name:          "Array with one element (element absent)",
			arr:           []int{1},
			query:         2,
			expectedIndex: -1,
		},
		{
			name:          "Duplicate elements in the array",
			arr:           []int{1, 2, 2, 2, 3, 4, 5},
			query:         2,
			expectedIndex: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearch(tc.arr, tc.query)

			if result != tc.expectedIndex {
				t.Errorf("Failed %s: expected %d, got %d", tc.name, tc.expectedIndex, result)
			}

			t.Logf("Success %s: expected %d, got %d", tc.name, tc.expectedIndex, result)
		})
	}
}

