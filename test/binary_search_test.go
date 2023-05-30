package test

import (
	"programmercarl/src/array"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	target := 3
	expected := 1
	result := array.Search(nums, target)
	if result != expected {
		t.Errorf("search(%v, %d) = %d; want %d", nums, target, result, expected)
	}

	nums = []int{-10, -5, 0, 1, 3, 5, 7, 9}
	target = 5
	expected = 5
	result = array.Search(nums, target)
	if result != expected {
		t.Errorf("search(%v, %d) = %d; want %d", nums, target, result, expected)
	}

	nums = []int{2, 4, 6, 8, 10}
	target = 7
	expected = -1
	result = array.Search(nums, target)
	if result != expected {
		t.Errorf("search(%v, %d) = %d; want %d", nums, target, result, expected)
	}
}
