package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721


 */
func TestlinearSearch(t *testing.T) {

	type testCase struct {
		description   string
		arr           []int
		query         int
		expectedIndex int
	}

	testCases := []testCase{
		{
			description:   "Query Present at Start",
			arr:           []int{5, 2, 3, 4, 1},
			query:         5,
			expectedIndex: 0,
		},
		{
			description:   "Query Present in the Middle",
			arr:           []int{1, 3, 5, 7, 9},
			query:         5,
			expectedIndex: 2,
		},
		{
			description:   "Query Present at End",
			arr:           []int{2, 3, 4, 5, 9},
			query:         9,
			expectedIndex: 4,
		},
		{
			description:   "Query Not Present",
			arr:           []int{1, 3, 5, 7, 9},
			query:         2,
			expectedIndex: -1,
		},
		{
			description:   "Empty Array",
			arr:           []int{},
			query:         5,
			expectedIndex: -1,
		},
		{
			description:   "All Elements the Same But Different From Query",
			arr:           []int{1, 1, 1, 1},
			query:         2,
			expectedIndex: -1,
		},
		{
			description:   "Multiple Occurrences of the Query",
			arr:           []int{2, 3, 5, 5, 5, 9},
			query:         5,
			expectedIndex: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Running scenario: %s", tc.description)
			result := linearSearch(tc.arr, tc.query)

			if result != tc.expectedIndex {
				t.Errorf("Failed: %s. Expected index %d, got %d", tc.description, tc.expectedIndex, result)
			} else {
				t.Logf("Success: %s. Query found at index %d as expected", tc.description, result)
			}
		})
	}

}

