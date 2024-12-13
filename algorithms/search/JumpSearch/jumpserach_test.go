package JumpSearch

import "testing"

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95


 */
func TestjumpSearch(t *testing.T) {

	type testCase struct {
		name     string
		arr      []int
		query    int
		expected int
	}

	tests := []testCase{
		{
			name:     "Scenario 1: Search for an Existing Element",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
			query:    7,
			expected: 3,
		},
		{
			name:     "Scenario 2: Search for a Non-Existent Element",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
			query:    4,
			expected: -1,
		},
		{
			name:     "Scenario 3: Search in an Empty Array",
			arr:      []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "Scenario 4: Search for the First Element",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
			query:    1,
			expected: 0,
		},
		{
			name:     "Scenario 5: Search for the Last Element",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
			query:    19,
			expected: 9,
		},
		{
			name:     "Scenario 6: Search with Step Exactly Matching Array Length",
			arr:      []int{10},
			query:    11,
			expected: -1,
		},
	}

	for _, tc := range tests {

		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			result := jumpSearch(tc.arr, tc.query)

			if result != tc.expected {
				t.Errorf("Test %s failed: expected %d but got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed: expected %d and got %d", tc.name, tc.expected, result)
			}
		})
	}

}

