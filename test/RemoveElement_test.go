package test

import (
	"programmercarl/src/array"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	// Test case 1: Normal input
	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	expected1 := []int{2, 2}
	expectedLength1 := 2
	actualLength1 := array.RemoveElement(nums1, val1)
	if actualLength1 != expectedLength1 || !reflect.DeepEqual(nums1[:actualLength1], expected1) {
		t.Errorf("Test case 1 failed: expected %v and %v, but got %v and %v", expected1, expectedLength1, nums1[:actualLength1], actualLength1)
	}

	// Test case 2: Empty input
	nums2 := []int{}
	val2 := 0
	expected2 := []int{}
	expectedLength2 := 0
	actualLength2 := array.RemoveElement(nums2, val2)
	if actualLength2 != expectedLength2 || !reflect.DeepEqual(nums2[:actualLength2], expected2) {
		t.Errorf("Test case 2 failed: expected %v and %v, but got %v and %v", expected2, expectedLength2, nums2[:actualLength2], actualLength2)
	}

	//Test case 3: Input with all elements equal to val
	nums3 := []int{4, 4, 4, 4}
	val3 := 4
	expected3 := []int{}
	expectedLength3 := 0
	actualLength3 := array.RemoveElement(nums3, val3)
	if actualLength3 != expectedLength3 || !reflect.DeepEqual(nums3[:actualLength3], expected3) {
		t.Errorf("Test case 3 failed: expected %v and %v, but got %v and %v", expected3, expectedLength3, nums3[:actualLength3], actualLength3)
	}

	// Test case 4: Input with no elements equal to val
	nums4 := []int{1, 2, 3, 4}
	val4 := 5
	expected4 := []int{1, 2, 3, 4}
	expectedLength4 := 4
	actualLength4 := array.RemoveElement(nums4, val4)
	if actualLength4 != expectedLength4 || !reflect.DeepEqual(nums4[:actualLength4], expected4) {
		t.Errorf("Test case 4 failed: expected %v and %v, but got %v and %v", expected4, expectedLength4, nums4[:actualLength4], actualLength4)
	}

	// Test case 5: Input with all elements equal to val except for the last one
	nums5 := []int{2, 2, 2, 2, 3}
	val5 := 2
	expected5 := []int{3}
	expectedLength5 := 1
	actualLength5 := array.RemoveElement(nums5, val5)
	if actualLength5 != expectedLength5 || !reflect.DeepEqual(nums5[:actualLength5], expected5) {
		t.Errorf("Test case 5 failed: expected %v and %v, but got %v and %v", expected5, expectedLength5, nums5[:actualLength5], actualLength5)
	}
}
