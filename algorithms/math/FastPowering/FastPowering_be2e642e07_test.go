// Test generated by RoostGPT for test go-unit-algo-string using AI Type Open AI and AI Model gpt-4

package FastPowering

import (
	"testing"
	"math"
)

func TestFastPowering_be2e642e07(t *testing.T) {

	// Test Case 1: Testing with base as 2 and power as 3. Expected result is 8
	result := fastPowering(2, 3)
	if result != 8 {
		t.Error("Test Case 1 Failed: Expected 8, got ", result)
	} else {
		t.Log("Test Case 1 Passed: Expected 8, got ", result)
	}

	// Test Case 2: Testing with base as 2 and power as 0. Expected result is 1
	result = fastPowering(2, 0)
	if result != 1 {
		t.Error("Test Case 2 Failed: Expected 1, got ", result)
	} else {
		t.Log("Test Case 2 Passed: Expected 1, got ", result)
	}

	// Test Case 3: Testing with base as 1 and power as 100. Expected result is 1
	result = fastPowering(1, 100)
	if result != 1 {
		t.Error("Test Case 3 Failed: Expected 1, got ", result)
	} else {
		t.Log("Test Case 3 Passed: Expected 1, got ", result)
	}

	// Test Case 4: Testing with base as 0 and power as 5. Expected result is 0
	result = fastPowering(0, 5)
	if result != 0 {
		t.Error("Test Case 4 Failed: Expected 0, got ", result)
	} else {
		t.Log("Test Case 4 Passed: Expected 0, got ", result)
	}

	// Test Case 5: Testing with base as -2 and power as 3. Expected result is -8
	result = fastPowering(-2, 3)
	if result != -8 {
		t.Error("Test Case 5 Failed: Expected -8, got ", result)
	} else {
		t.Log("Test Case 5 Passed: Expected -8, got ", result)
	}

	// Test Case 6: Testing with base as 2 and power as -3. Expected result is 0.125
	result = fastPowering(2, -3)
	expected := 1 / math.Pow(2, 3)
	if result != expected {
		t.Error("Test Case 6 Failed: Expected", expected, "got ", result)
	} else {
		t.Log("Test Case 6 Passed: Expected", expected, "got ", result)
	}
}
