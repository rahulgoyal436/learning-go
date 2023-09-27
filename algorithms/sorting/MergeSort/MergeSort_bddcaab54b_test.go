// Test generated by RoostGPT for test roost-test using AI Type Open AI and AI Model gpt-4

package MergeSort

import (
	"reflect"
	"testing"
)

func TestMergeSort_bddcaab54b(t *testing.T) {
	// Test case 1: Normal scenario
	arr1 := []int{3, 2, 1}
	expected1 := []int{1, 2, 3}
	result1 := mergeSort(arr1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected1, result1)
	} else {
		t.Logf("TestMergeSort_bddcaab54b passed for input %v", arr1)
	}

	// Test case 2: Edge case, empty array
	arr2 := []int{}
	expected2 := []int{}
	result2 := mergeSort(arr2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected2, result2)
	} else {
		t.Logf("TestMergeSort_bddcaab54b passed for input %v", arr2)
	}

	// Test case 3: Array with one element
	arr3 := []int{1}
	expected3 := []int{1}
	result3 := mergeSort(arr3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected3, result3)
	} else {
		t.Logf("TestMergeSort_bddcaab54b passed for input %v", arr3)
	}

	// Test case 4: Array with duplicate elements
	arr4 := []int{5, 3, 5, 1, 2, 3}
	expected4 := []int{1, 2, 3, 3, 5, 5}
	result4 := mergeSort(arr4)
	if !reflect.DeepEqual(result4, expected4) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected4, result4)
	} else {
		t.Logf("TestMergeSort_bddcaab54b passed for input %v", arr4)
	}
}
