package main

import "testing"

func TestThreeSum(t *testing.T) {
	checkThreeSum(t, 6, []int{1, 2, 3}, true)
	checkThreeSum(t, 3, []int{1, 4, 3}, false)
	checkThreeSum(t, 4, []int{1, 2, 1, 0, 4}, true)
	checkThreeSum(t, 4, []int{0, 2, 1, 4, 4}, false)
	checkThreeSum(t, 4, []int{-1, 2, 5}, false)
	checkThreeSum(t, 6, []int{-1, 2, 5}, true)
	checkThreeSum(t, 14, []int{2, 12, 3, 0, 3, 5, 9}, true)
	checkThreeSum(t, 14, []int{1, 1, 2, 3, 2, 10, 4}, true)
}

func checkThreeSum(t *testing.T, s int, a []int, e bool) {
	r := ThreeSum(s, a)
	if r != e {
		t.Errorf("Expected %t from (%d, %d) but got %t", e, s, a, r)
	}
}

func ThreeSum(t int, a []int) bool {
	for i, v := range a {
		ai := append(a[:i], a[i+1:]...)
		if TwoSum(t-v, ai) {
			return true
		}
	}
	return false
}
